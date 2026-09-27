package main

import (
	"fmt"
	funct "groupie-tracker/functions"
	"groupie-tracker/structdata"
	"net/http"
)

func main() {
	err := structdata.LoadData()
	if err != nil {
		fmt.Println(err)
		return
	}
	funct.LoadTemplates()
	http.HandleFunc("/", funct.HomePage)
	http.HandleFunc("/artist", funct.ArtistPage)
	http.HandleFunc("/search", funct.Search)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("picture"))))
	fmt.Println("Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
