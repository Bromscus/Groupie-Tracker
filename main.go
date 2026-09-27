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
	http.HandleFunc("/", funct.ProtectionForHandler(funct.HomePage))
	http.HandleFunc("/artist", funct.ProtectionForHandler(funct.ArtistPage))
	http.HandleFunc("/search", funct.ProtectionForHandler(funct.Search))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/pict/", http.StripPrefix("/pict/", http.FileServer(http.Dir("pict"))))
	fmt.Println("Server running at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
