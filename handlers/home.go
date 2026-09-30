package functions

import (
	"net/http"
)

// templ.Execute basically means we write the final HTML into the HTTP response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, E404, E404Msg, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		ErrorPage(w, E405, E405HomeMsg, http.StatusMethodNotAllowed)
		return
	}

	renderArtists(w, Artists)
}
