package infra

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template is a type for html/template.
type Template struct {
	templates *template.Template
}

// NewTemplate creates a new Template using globbed files.
func NewTemplate(glob string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(glob)),
	}
}

// Render processes template rendering.
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
