package main

import (
	"fmt"
	functions "groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	err := functions.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", functions.HomeHandler)
	http.HandleFunc("/search", functions.SearchHandler)
	http.HandleFunc("/artist", functions.ArtistHandler)

	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
