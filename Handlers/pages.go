package functions

import (
	"groupie-tracker/structdata"
	"net/http"
	"strconv"
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
