package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/kazgorot/echo_oapi/docs" // Import the docs
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/api/v1/test1", test1)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce text/plain
// @Success 200 {object} string
// @Router / [get]
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// HealthCheck test1
// @Summary Show the test1.
// @Description get the status of test1.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/test1 [get]
func test1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, test1!")
}
