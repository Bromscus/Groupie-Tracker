package main

import (
	"fmt"
	functions "groupie-tracker/Handlers"
	"groupie-tracker/structdata"
	"net/http"
)

func main() {
	err := structdata.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}
	functions.LoadTemplates()
	http.HandleFunc("/", functions.HomePage)
	http.HandleFunc("/artist", functions.ArtistPage)
	http.HandleFunc("/search", functions.Search)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("picture"))))
	fmt.Println("Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
