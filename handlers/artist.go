package functions

import (
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorPage(w, E404, E404Msg, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPage(w, E405, E405ArtistMsg, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		ErrorPage(w, E400, E400ArtistMsg, http.StatusBadRequest)
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
		ErrorPage(w, E404Artist, E404ArtistMsg, http.StatusNotFound)
		return
	}

	templ, err := template.ParseFiles("templates/info.html")
	if err != nil {
		ErrorPage(w, E500, E500Msg, http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, selectedArtist); err != nil {
		ErrorPage(w, E500, E500Msg, http.StatusInternalServerError)
	}
}

func renderArtists(w http.ResponseWriter, artists []Artist) {
	templ, err := template.ParseFiles("templates/homepage.html")
	if err != nil {
		ErrorPage(w, E500, E500Msg, http.StatusInternalServerError)
		return
	}

	if err := templ.Execute(w, artists); err != nil {
		ErrorPage(w, E500, E500Msg, http.StatusInternalServerError)
	}
}
