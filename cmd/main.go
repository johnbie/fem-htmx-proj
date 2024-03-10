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

func (d *Data) hasEmail(email string) bool {
    for _, contact := range d.Contacts {
        if contact.Email == email {
            return true
        }
    }
    return false
}

func newData() Data {
    return Data {
        Contacts: []Contact {
            newContact("John", "jd@gmail.com"),
            newContact("Clara", "cd@gmail.com"),
        },
    }
}

type FormData struct {
    Values map[string]string
    Errors map[string]string
}

func newFormData() FormData {
   return FormData {
        Values: make(map[string]string),
	Errors: make(map[string]string),
   }
}

type Page struct {
    Data Data
    Form FormData
}

func newPage() Page {
    return Page {
        Data: newData(),
        Form: newFormData(),
    }
}

func main() {

    e := echo.New()

    page := newPage()

    e.Renderer = newTemplate()
    e.Use(middleware.Logger())

    e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", page)
    });

    e.POST("/contacts", func(c echo.Context) error {
        name := c.FormValue("name")
	email := c.FormValue("email")

        if page.Data.hasEmail(email) {
            formData := newFormData()
            formData.Values["name"] = name
            formData.Values["email"] = email
	    formData.Errors["email"] = "Email already exists"
            return c.Render(422, "form", formData)
        }
	page.Data.Contacts = append(page.Data.Contacts, newContact(name, email))

	return c.Render(200, "display", page)
    });

    e.Logger.Fatal(e.Start(":42069"))
}
