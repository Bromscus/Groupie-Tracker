package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "405: Method Not Allowed", "Artist information is available with a GET request.", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		ErrorPage(w, "400: Bad Request", "Please provide a valid artist ID.", http.StatusBadRequest)
		return
	}

	var selectedArtist Artist

	for _, artist := range Artists {
		if artist.ID == id {
			selectedArtist = artist
			break
		}
	}

	if selectedArtist.ID == 0 {
		ErrorPage(w, "404: Artist Not Found", "No artist exists with that ID.", http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("templates/info.html")
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, selectedArtist); err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
	}
}

func renderArtists(w http.ResponseWriter, artists []Artist) {
	templ, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, artists); err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
	}
}
