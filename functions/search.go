package funct

import (
	"groupie-tracker/structdata"
	"net/http"
	"strings"
)

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
