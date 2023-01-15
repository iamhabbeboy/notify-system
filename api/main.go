package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	// integration
	alert := NewAlert()
	alert.Server(":1323")
	alert.SetNotificationService("email", "slack", "telegram")
	// You have to add each handler for each service.

	e := echo.New()
	e.HTTPErrorHandler = CustomHTTPErrorHandler
	e.GET("/", func(c echo.Context) error {
		t := time.Now().Format(time.RFC3339)
		// logging
		alert.Log(fmt.Sprintf("LOG:[%v] error occured here this is a sample for error logger", t))
		return c.String(http.StatusOK, "alert-notify is integerated with the endpoint !")
	})

	e.Logger.Fatal(e.Start(":9000"))
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	// if er, ok := err.(*echo.HTTPError); ok {
	// 	// logger.Log(string(er.Error()))
	// 	c.Logger().Error(err)
	// }
}
