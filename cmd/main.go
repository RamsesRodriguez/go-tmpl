package main

import (
	"go-tmpl/handlers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

var GO_ENV = os.Getenv("GO_ENV")

// @title Go clean echo API v1
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8888
// @BasePath /
// @schemes http
func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	helloHandler := handlers.HelloHandler{}
	e.GET("/", helloHandler.HandlerHelloShow)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}

	e.Logger.Fatal("failed to start server", zap.Error(e.Start(":"+PORT)))
}
