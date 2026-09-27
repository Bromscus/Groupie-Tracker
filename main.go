package main

import (
	"fmt"
	handlers "groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	// authors image, location, dates, relations all have link
	http.Handle(
		"/style/",
		http.StripPrefix("/style/", http.FileServer(http.Dir("style"))),
	)

	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
