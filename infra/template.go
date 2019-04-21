package infra

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template は html/template を使ったテンプレート型。
type Template struct {
	templates *template.Template
}

// NewTemplate は glob のファイル群を使って新しい Template を生成する。
func NewTemplate(glob string) *Template {
	return &Template{
		templates: template.Must(template.ParseGlob(glob)),
	}
}

// Render はテンプレートの処理を行う。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
