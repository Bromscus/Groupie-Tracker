package functions

import (
	"encoding/json"
	"log"
	"net/http"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func reader() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: status code %d", response.StatusCode)
	}
	api := API{}

	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&api); err != nil {
		log.Fatal(err)
	}

	artistsResponse, err := http.Get(api.Artists)
	if err != nil {
		log.Fatal(err)
	}
	defer artistsResponse.Body.Close()

	if artistsResponse.StatusCode != http.StatusOK {
		log.Fatal("failed to fetch artists")
	}

	var artists []Artist

	decoder = json.NewDecoder(artistsResponse.Body)

	if err := decoder.Decode(&artists); err != nil {
		log.Fatal(err)
	}

	return artists, nil
}
