package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Error   string
	Message string
}

const (
	E400          = "400: Bad Request"
	E400ArtistMsg = "Please provide a valid artist ID."
	E404          = "404: page not found"
	E404Msg       = "Please check that the Website address is spelled correctly or go to our home page."
	E404Artist    = "404: Artist Not Found"
	E404ArtistMsg = "No artist exists with that ID."
	E405          = "405: Method Not Allowed"
	E405ArtistMsg = "Artist information is available with a GET request."
	E405HomeMsg   = "This page is available with a GET request."
	E405SearchMsg = "Search is available with a GET request."
	E500          = "500: Internal Server Error"
	E500Msg       = "Something went wrong on the server."
)

func ErrorPage(w http.ResponseWriter, errorText string, msg string, status int) {
	templ, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, msg, status)
		return
	}

	w.WriteHeader(status)
	data := ErrorData{
		Error:   errorText,
		Message: msg,
	}
	err = templ.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
