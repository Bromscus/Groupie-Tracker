package main

import (
	"fmt"
	handler "groupie-tracker/Handlers"
	"log"
	"net/http"
)

func main() {
	// authors image, location, dates, relations all have link

	http.HandleFunc("/", handler.HomeHandler)
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
