package handler

import (
	"go-tmpl/views/hello"

	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h HelloHandler) HandlerHelloShow(c echo.Context) error {
	return render(c, hello.ShowHello("Ramses"))
}
