<div align="center">
  <a href="https://github.com/S6-BikePack">
    <img src="assets/logo.png" alt="logo" width="200" height="auto" />
  </a>
  <h1>BikePack - Authentication-Service</h1>

  <p>
    Part of the S6 BikePack project.
  </p>


<!-- Badges -->
<p>

</p>

<h4>
    <a href="https://github.com/S6-BikePack">Home</a>
  <span> · </span>
    <a href="https://github.com/S6-BikePack/authentication-service#-about-the-project">Documentation</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->
# 📓 Table of Contents

- [About the Project](#-about-the-project)
    * [Tech Stack](#%EF%B8%8F-tech-stack)
    * [Environment Variables](#-environment-variables)
- [Getting Started](%EF%B8%8F-getting-started)
    * [Prerequisites](%EF%B8%8F-prerequisites)
    * [Running Tests](#-running-tests)
    * [Run Locally](#-run-locally)
    * [Deployment](#-deployment)
- [Usage](#-usage)



<!-- About the Project -->
## ⭐ About the Project

The Authentication-Service is the service for the BikePack project that handles the authentication of users accessing the API's. 
When a request comes in at the api-gateway for one of the authenticated api's the gateway will forward the request to this service for authentication.


<!-- TechStack -->
### 🛰️ Tech Stack
#### Language
  <ul>
    <li><a href="https://go.dev/">GoLang</a></li>
</ul>

#### Dependencies
  <ul>
    <li><a href="https://github.com/gin-gonic/gin">Gin</a><span> - Web framework</span></li>
    <li><a href="https://github.com/swaggo/swag">Swag</a><span> - Swagger documentation</span></li>
    <li><a href="https://gorm.io/index.html">GORM</a><span> - ORM library</span></li>
  </ul>

<!-- Env Variables -->
### 🔑 Environment Variables

This service has the following environment variables that can be set:

`PORT` - Port the service runs on

<!-- Getting Started -->
## 	🛠️ Getting Started

<!-- Prerequisites -->
### ‼️ Prerequisites

Building the project requires Go 1.18.

Running the service requires a firebase serviceKey.json file in the root folder of the container.

The easiest way to setup the project is to use the Docker-Compose file from the infrastructure repository.

<!-- Running Tests -->
### 🧪 Running Tests

-

<!-- Run Locally -->
### 🏃 Run Locally

Clone the project

```bash
  git clone https://github.com/S6-BikePack/authentication-service
```

Go to the project directory

```bash
  cd authentication-service
```

Run the project (Rest)

```bash
  go run cmd/rest/main.go
```


<!-- Deployment -->
### 🚀 Deployment

To build this project run (Rest)

```bash
  go build cmd/rest/main.go
```


<!-- Usage -->
## 👀 Usage

### REST
Once the service is running you can find its swagger documentation with all the endpoints at `/swagger`