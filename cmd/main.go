package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const defaultPort = ":1234"

type Server struct {
	auth *auth.Client
}

func main() {
	var opt option.ClientOption
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") != "" {
		opt = option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	} else {
		opt = option.WithCredentialsFile("./serviceKey.json")
	}

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("Error initializing firebase: %v\n", err)
	}

	authentication, err := app.Auth(context.Background())
	if err != nil {
		fmt.Printf("Error initializing firebase authentication: %v\n", err)
	}

	server := Server{auth: authentication}

	router := gin.New()

	router.GET("/auth", server.Authorize)

	port := GetEnvOrDefault("PORT", defaultPort)

	log.Fatal(router.Run(port))
}

func (s *Server) Authorize(c *gin.Context) {
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

	c.AbortWithStatus(200)
}

func GetEnvOrDefault(environmentKey, defaultValue string) string {
	returnValue := os.Getenv(environmentKey)
	if returnValue == "" {
		returnValue = defaultValue
	}
	return returnValue
}
