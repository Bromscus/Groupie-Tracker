package funct

import (
	"fmt"
	"groupie-tracker/errorpage"
	"groupie-tracker/structdata"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/" || r.Method != "GET" {
		errorpage.RenderError(w, "Page not found", http.StatusNotFound)
		return
	}
	
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errorpage.RenderError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = templ.Execute(w, structdata.Artists)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorpage.RenderError(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	idNum, err := strconv.Atoi(idStr)
	if err != nil {
		errorpage.RenderError(w, "invalid id", http.StatusBadRequest)
		return
	}

	
	var detail structdata.ArtistDetails
	foundArtist := false
	foundLocation := false
	foundDate := false
	foundRelation := false 
	for _, artist := range structdata.Artists {
		if artist.ID == idNum {
			detail.Artist = artist
			foundArtist = true
			break
		}

	}
	for _, location := range structdata.Locations {
		if location.ID == idNum {
			detail.Location = location
			foundLocation = true
			break
		}

	}
	for _, date := range structdata.Dates {
		if date.ID == idNum {
			detail.Date = date
			foundDate = true
			break
		}

	}
	for _, relation := range structdata.Relations {
		if relation.ID == idNum {
			detail.Relation = relation
			foundRelation = true
			break
		}

	}
	if !foundArtist  {
		errorpage.RenderError(w, "artist data not found", http.StatusNotFound)
		return
	}
	if !foundLocation  {
		errorpage.RenderError(w, "location data not found", http.StatusNotFound)
		return
	}
	if !foundDate  {
		errorpage.RenderError(w, "date data not found", http.StatusNotFound)
		return
	}
	if !foundRelation  {
		errorpage.RenderError(w, "relation data not found", http.StatusNotFound)
		return
	}




	tmple, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w,"error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmple.Execute(w, detail)
	if err != nil {
		http.Error(w,"error", http.StatusInternalServerError)
		return
	}
	
}

func ProtectionForHandler(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {
		defer func(){
			if err := recover(); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		}()
		next(w,r)
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	matchedArtists := []structdata.Artist{}
	artistName := r.URL.Query().Get("q")
	artistName = strings.TrimSpace(artistName)

	for _, artist := range structdata.Artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(artistName)) {
			matchedArtists = append(matchedArtists, artist)
		}
	}

	tmple, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmple.Execute(w, matchedArtists)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return 
	}
}