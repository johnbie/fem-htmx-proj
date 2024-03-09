package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type Templates struct {
	template *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecutTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),

	}
}

func main() {

	e := echo.New()
	count := Count { Count: 0}

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "index", count)
	});

	e.Logger.Fatal(e.Start(":1234"))
}
