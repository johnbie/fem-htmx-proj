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
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

}
