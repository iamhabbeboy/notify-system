package main

import (
	"bufio"
	"fmt"
	"main/cmd/service"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var logger = service.NewLoggerService()

func main() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")

	logger.SetLogger("slack", "telegram", "email")
	e := echo.New()
	UseCors(e)
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	e.GET("/", func(c echo.Context) error {
		t := time.Now().Format(time.RFC3339)
		logger.Log(fmt.Sprintf("LOG:[%v] error occured here this is a sample for error logger", t))
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/logs", handleLogs)
	e.GET("/setting", handleSettings)

	e.Logger.Fatal(e.Start(":1323"))
}

func UseCors(router *echo.Echo) {
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
}

func handleLogs(c echo.Context) error {
	f, err := os.Open("local.log")
	if err != nil {
		resp := map[string]string{"status": "false", "data": err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	output := []map[string]string{}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		payload := fileScanner.Text()
		if payload == "" {
			continue
		}
		splitter := strings.Split(payload, "::")
		o := map[string]string{
			"serviceName": strings.Trim(splitter[0], "[]"),
			"environment": strings.Trim(splitter[1], "[]"),
			"date":        strings.Trim(splitter[2], "[]"),
			"message":     strings.Trim(splitter[3], " "),
		}
		output = append(output, o)
	}

	result := make(map[string]interface{})
	result["status"] = "true"
	result["data"] = output

	return c.JSON(http.StatusOK, result)
}

func handleSettings(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Setting!")
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if er, ok := err.(*echo.HTTPError); ok {
		logger.Log(string(er.Error()))
		c.Logger().Error(err)
	}
}
