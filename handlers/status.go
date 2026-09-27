package functions

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// templ.Execute basically means we write the final HTML into the HTTP response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)

		return
	}
	artists, err := reader()
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	renderArtists(w, artists)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "405: Method Not Allowed", "Search is available with a GET request.", http.StatusMethodNotAllowed)
		return
	}

	artists, err := reader()
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))
	filteredArtists := make([]Artist, 0, len(artists))
	for _, artist := range artists {
		if artistMatchesSearch(artist, query) {
			filteredArtists = append(filteredArtists, artist)
		}
	}
	renderArtists(w, filteredArtists)
}

func artistMatchesSearch(artist Artist, query string) bool {
	if query == "" {
		return true
	}
	matches := func(value string) bool {
		return strings.Contains(strings.ToLower(value), query)
	}
	if matches(artist.Name) || matches(strconv.Itoa(artist.CreationDate)) || matches(artist.FirstAlbum) {
		return true
	}
	for _, member := range artist.Members {
		if matches(member) {
			return true
		}
	}
	for _, location := range artist.Locations {
		if matches(location) {
			return true
		}
	}
	for _, date := range artist.Dates {
		if matches(date) {
			return true
		}
	}
	for location, dates := range artist.Relations {
		if matches(location) {
			return true
		}
		for _, date := range dates {
			if matches(date) {
				return true
			}
		}
	}
	return false
}

func renderArtists(w http.ResponseWriter, artists []Artist) {
	templ, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	err = templ.Execute(w, artists)
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
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

	artists, err := reader()
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}

	var selectedArtist Artist
	for _, artist := range artists {
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
