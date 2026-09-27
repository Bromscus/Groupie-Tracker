package errorpage

import (
	"html/template"
	"net/http"
)

type ErrorDetail struct {
	Code    int
	Message string
}

func RenderError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, ErrorDetail{Message: message, Code: code})
}