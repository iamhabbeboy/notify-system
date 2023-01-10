package main

import (
	"fmt"
	"main/cmd/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

var logger = services.NewLoggerService()

func main() {
	logger.SetLogger("slack", "telegram", "email")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.HTTPErrorHandler = CustomHTTPErrorHandler
	e.Logger.Fatal(e.Start(":1323"))
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if er, ok := err.(*echo.HTTPError); ok {
		logger.Log(string(er.Error()))
		var pcode string
		switch er.Code {
		case http.StatusNotFound:
			pcode = "404"
		case http.StatusUnauthorized:
			c.String(http.StatusUnauthorized, er.Message.(string))
		}
		c.Render(http.StatusOK, fmt.Sprintf("%s.html", pcode), nil)
	}
}
