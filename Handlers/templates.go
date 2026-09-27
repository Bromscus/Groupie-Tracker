package functions

import (
	"fmt"
	"html/template"
)


var templates map[string]*template.Template

func LoadTemplates() {
	templates = make(map[string]*template.Template)
	for _, path := range []string{"templates/index.html", "templates/artist.html", "templates/error.html"} {
		templ, err := template.ParseFiles(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		templates[path] = templ
	}
}
