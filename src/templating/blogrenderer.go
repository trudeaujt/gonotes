package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//With embedding, the template files are included into our Go program when we build it.
//This means once we've built it, the files are always available, we don't have to worry about dragging them around with our program!
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.Execute(w, p); err != nil {
		return err
	}
	return nil
}
