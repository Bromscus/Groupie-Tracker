package structdata

import (
	"encoding/json"
	"net/http"
	"strings"
)

var (
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
)

type Artist struct {
	Name         string
	ID           int
	Members      []string
	Image        string
	FirstAlbum   string
	CreationDate int
}

type Location struct {
	ID        int
	Locations []string
}

type LocationWrapper struct {
	Index []Location
}

type Date struct {
	ID    int
	Dates []string
}

type DateWrapper struct {
	Index []Date
}

type Relation struct {
	ID             int
	DatesLocations map[string][]string
}

type RelationWrapper struct {
	Index []Relation
}

type ArtistDetails struct {
	Artist   Artist
	Location Location
	Relation Relation
	Date     Date
}

func FetchData(url string, destination any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(destination)
}

func LoadData() error {
	if err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &Artists); err != nil {
		return err
	}
	for i := range Artists {
		name := strings.TrimSpace(Artists[i].Name)
		if strings.EqualFold(name, "Mamonas Assassinas") || strings.EqualFold(name, "Mamonas Assasinans") {
			Artists[i].Image = "/picture/mamonas.jpg"
		}
	}

	var locations LocationWrapper
	if err := FetchData("https://groupietrackers.herokuapp.com/api/locations", &locations); err != nil {
		return err
	}
	Locations = locations.Index

	var dates DateWrapper
	if err := FetchData("https://groupietrackers.herokuapp.com/api/dates", &dates); err != nil {
		return err
	}
	Dates = dates.Index

	var relations RelationWrapper
	if err := FetchData("https://groupietrackers.herokuapp.com/api/relation", &relations); err != nil {
		return err
	}
	Relations = relations.Index

	buildIndexes()
	return nil
}
