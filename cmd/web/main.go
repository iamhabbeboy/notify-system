package main

import (
	"fmt"
	"main/cmd/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var logger = service.NewLoggerService()

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")

	logger.SetLogger("slack", "telegram", "email")
	e := echo.New()
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	e.GET("/", func(c echo.Context) error {
		t := time.Now().Format(time.RFC3339)
		logger.Log(fmt.Sprintf("LOG:[%v] error occured here this is a sample for error logger", t))
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if er, ok := err.(*echo.HTTPError); ok {
		logger.Log(string(er.Error()))
		c.Logger().Error(err)
	}
}
