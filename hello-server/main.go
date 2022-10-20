package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello/:name", helloHandler)

	e.Start(":8080")
}

func helloHandler(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+".\n")
}
