package funct

import (
	"bytes"
	"fmt"
	"net/http"
)

type ErrorData struct {
	Status  int
	Error   string
	Message string
}

func ErrorPage(w http.ResponseWriter, errorText string, message string, status int) {
	templ := templates["templates/error.html"]
	if templ == nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := ErrorData{
		Status:  status,
		Error:   errorText,
		Message: message,
	}

	var body bytes.Buffer
	if err := templ.Execute(&body, data); err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(body.Bytes()); err != nil {
		fmt.Println(err)
	}
}
