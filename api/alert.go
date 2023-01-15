package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Alert struct {
	logger Logger
}

func NewAlert() *Alert {
	getConfig()

	return &Alert{
		logger: *NewLogger(),
	}
}

func (a *Alert) Server(port string) {
	e := echo.New()
	UseCors(e)

	e.GET("/logs", handleLogs)
	e.GET("/setting", handleSettings)

	go func() {
		e.Logger.Fatal(e.Start(port))
	}()
}

func (a *Alert) SetNotificationService(data ...string) {
	a.logger.SetLogger(data...)
}

func (a *Alert) Log(data string) {
	a.logger.Log(data)
}

func getConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	var sample = []byte(`
appName: notify coffee
environment: production
services:
- slack:
  - webhook: http://
- email:
  - host: smtp
`)
	viper.ReadConfig(bytes.NewBuffer(sample))
	viper.SafeWriteConfig()
	cnf := map[string]string{
		"name": "abiodun",
		"age":  "14",
	}

	sm := reflect.ValueOf(cnf).MapKeys()
	fmt.Println(sm)
	a := viper.Get("services")
	fmt.Println(a)
	// logger.SetLogger("slack", "telegram", "email")
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
