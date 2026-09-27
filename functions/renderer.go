package funct

import (
	"bytes"
	"fmt"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, path string, data any) {
	templ := templates[path]
	if templ == nil {
		ErrorPage(w, "Internal server error", "The page could not be loaded.", http.StatusInternalServerError)
		return
	}

	var body bytes.Buffer

	if err := templ.Execute(&body, data); err != nil {
		ErrorPage(w, "Internal server error", "The page could not be rendered.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if _, err := w.Write(body.Bytes()); err != nil {
		fmt.Println(err)
	}
}
