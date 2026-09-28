package functions

import (
	"net/http"
)

// templ.Execute basically means we write the final HTML into the HTTP response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		ErrorPage(w, "405: Method Not Allowed", "This page is available with a GET request.", http.StatusMethodNotAllowed)
		return
	}

	renderArtists(w, Artists)
}
