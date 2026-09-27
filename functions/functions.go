package funct

import (
	"bytes"
	"fmt"
	"groupie-tracker/structdata"
	"net/http"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "Page not found", "Please check that the website address is spelled correctly or go to our home page.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "Method not allowed", "This page only supports GET requests.", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "templates/index.html", structdata.Artists)
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "Method not allowed", "This page only supports GET requests.", http.StatusMethodNotAllowed)
		return
	}

	idNum, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		ErrorPage(w, "Invalid artist ID", "Please provide a valid artist ID.", http.StatusBadRequest)
		return
	}

	detail, found := structdata.GetArtistDetails(idNum)
	if !found {
		ErrorPage(w, "Artist not found", "No complete artist data exists with that ID.", http.StatusNotFound)
		return
	}

	renderTemplate(w, "templates/artist.html", detail)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "Method not allowed", "This page only supports GET requests.", http.StatusMethodNotAllowed)
		return
	}

	search := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

	matchedArtists := structdata.SearchArtists(search)

	renderTemplate(w, "templates/index.html", matchedArtists)
}

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
