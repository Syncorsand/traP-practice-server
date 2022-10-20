package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jsonData struct {
	Number int
	String string
	Bool   bool
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// e.GET("/hello", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello,World\n")
	// })

	e.POST("/post", postHandler)

	// e.GET("/json", jsonHandler)

	e.Start(":8080")
}

// func jsonHandler(c echo.Context) error {
// 	res := jsonData{
// 		Number: 10,
// 		String: "hoge",
// 		Bool:   false,
// 	}

// 	return c.JSON(http.StatusOK, &res)
// }

func postHandler(c echo.Context) error {
	data := jsonData{
		Number: 30,
		String: "huge",
		Bool:   true,
	}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}
