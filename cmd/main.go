package main

import (
	"go-tmpl/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	helloHandler := handler.HelloHandler{}

	app.GET("/", helloHandler.HandlerHelloShow)

	app.Start(":3000")
}
