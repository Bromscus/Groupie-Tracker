package functions

import (
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorPage(w, E404, E404Msg, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPage(w, E405, E405SearchMsg, http.StatusMethodNotAllowed)
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

	if matches(artist.Name) {
		return true
	}
	return false
}
