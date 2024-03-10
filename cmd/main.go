package main

import (
    "html/template"
    "io"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpl *template.Template
}

func newTemplate() *Template {
    return &Template{
        tmpl: template.Must(template.ParseGlob("views/*.html")),
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.tmpl.ExecuteTemplate(w, name, data)
}

type Count struct {
    Count int
}

type Contact struct {
    Name string
    Email string
}

func newContact(name, email string) Contact {
    return Contact {
        Name: name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func newData() Data {
    return Data {
        Contacts: []Contact {
            newContact("John", "jd@gmail.com"),
            newContact("Clara", "cd@gmail.com"),
        },
    }
}

func main() {

    e := echo.New()

    data := newData()

    e.Renderer = newTemplate()
    e.Use(middleware.Logger())

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    });

    e.POST("/contacts", func(c echo.Context) error {
        name := c.FormValue("name")
	email := c.FormValue("email")

        data.Contacts = append(data.Contacts, newContact(name, email))

	return c.Render(200, "display", data)
    });

    e.Logger.Fatal(e.Start(":42069"))
}
