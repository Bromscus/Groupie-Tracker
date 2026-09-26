package main

import (
	"fmt"
	funct "groupie-tracker/funct"
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
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
