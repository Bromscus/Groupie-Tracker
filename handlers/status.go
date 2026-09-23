package handler

import (
	valid "groupie-tracker/validation"
	"html/template"
	"net/http"
)

type authors struct {
	Artists []Artist
}

var (
	htmlHash = valid.HtmlHash()
	infoHash = valid.InfoHash()
)

// templ.Execute basically means we write the final HTML into the HTTP response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)

		return
	}
	// template allows Go to generate a HTML output
	if !valid.CheckFile("templates/homepage.html", htmlHash) {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	templ, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	artists, err := reader()
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	data := authors{
		Artists: artists,
	}

	err = templ.Execute(w, data)
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
}
