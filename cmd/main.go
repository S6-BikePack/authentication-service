package main

import (
	"authentication-service/pkg/tracing"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"strings"
)

const defaultPort = ":1234"

type Server struct {
	auth *auth.Client
}

func main() {
	tracer, err := tracing.TracerProvider("http://jaeger:6831/api/traces")

	if err != nil {
		panic(err)
	}

	otel.SetTracerProvider(tracer)

	opt := option.WithCredentialsFile("./serviceKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}

	authentication, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	server := Server{auth: authentication}

	router := gin.New()
	router.Use(otelgin.Middleware("authentication-service", otelgin.WithTracerProvider(tracer)))

	router.GET("/auth", server.Authorize)

	port := GetEnvOrDefault("PORT", defaultPort)

	log.Fatal(router.Run(port))
}

func (s *Server) Authorize(c *gin.Context) {
	span := trace.SpanFromContext(c)

	authorizationToken := c.GetHeader("Authorization")

	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "authorization token missing"})
		return
	}

	token, err := s.auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	user, err := s.auth.GetUser(context.Background(), token.UID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return
	}

	claims, _ := json.Marshal(user.CustomClaims)

	c.Header("x-user-claims", string(claims))
	c.Header("x-user-email", user.Email)
	c.Header("x-user-id", user.UID)

	span.End()
	c.AbortWithStatus(200)
}

func GetEnvOrDefault(environmentKey, defaultValue string) string {
	returnValue := os.Getenv(environmentKey)
	if returnValue == "" {
		returnValue = defaultValue
	}
	return returnValue
}
