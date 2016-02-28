package controllers

import (
	"net/http"
	"github.com/bmbstack/ripple"
	"github.com/labstack/echo"
)

func init() {
	ripple.RegisterController(&HomeController{})
}

type HomeController struct {
	Index echo.HandlerFunc `controller:"GET /"`
	Html echo.HandlerFunc `controller:"GET html"`
	String  echo.HandlerFunc `controller:"GET string"`
}

func (this HomeController) Path() string {
	return "/"
}

func (this HomeController) IndexFunc(ctx *echo.Context) error {
	ctx.Render(http.StatusOK, "home/index.html", map[string]interface{}{
		"title": "Hello, forum is a Ripple application ",
	})

	return nil
}

func (this HomeController) HtmlFunc(ctx *echo.Context) error {
	ctx.Render(http.StatusOK, "home/html.html", map[string]interface{}{
		"title": "Hello, this is a html template",
	})

	return nil
}

func (this HomeController) StringFunc(ctx *echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, this is a string")
}
