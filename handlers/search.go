package functions

import (
	"net/http"
	"strconv"
	"strings"
)

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

	query := strings.ToLower(strings.TrimSpace(r.URL.Query().Get("q")))

	filteredArtists := make([]Artist, 0, len(Artists))

	for _, artist := range Artists {
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
