package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

// integration
var alert = NewAlert()

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	alert.Server(":1323")
	alert.SetNotificationService("email", "slack", "telegram")
	// You have to add each handler for each service.

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}

	e.Renderer = renderer
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
	if er, ok := err.(*echo.HTTPError); ok {
		alert.Log(string(er.Error()))
		var pcode string
		switch er.Code {
		case http.StatusNotFound:
			pcode = "404"
		case http.StatusUnauthorized:
			c.String(http.StatusUnauthorized, er.Message.(string))
		}
		c.Render(http.StatusOK, fmt.Sprintf("%s.html", pcode), nil)

		c.Logger().Error(err)
	}
}
