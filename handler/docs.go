package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/maantje/postcode/view/example"
)

type DocHandler struct {
}

func (d DocHandler) Docs(c echo.Context) error {
	return render(c, example.Index())
}
