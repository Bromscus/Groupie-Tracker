package structdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var Artists []Artist
var Locations []Location
var Dates []Date
var Relations []Relation



type Artist struct {
	Name string 
	ID int 
	Members []string
	Image string 
	FirstAlbum string
	CreationDate int

}

type Location struct {
	ID int 
	Locations []string
}

type LocationWrapper struct {
	Index []Location
}

type Date struct {
	ID int 
	Dates []string
}

type DateWrapper struct {
	Index []Date
}

type Relation struct {
	ID int 
	DatesLocations map[string][]string
}

type RelationWrapper struct {
	Index []Relation
}


type ArtistDetails struct {
	Artist Artist
	Location Location
	Relation Relation
	Date Date
}

func FetchData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body

}

func UnmarshalsArtist(rawBytes []byte) ([]Artist, error) {
	var data []Artist

	err := json.Unmarshal(rawBytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UnmarshalsLocation(rawBytes []byte) ([]Location , error) {
	var data LocationWrapper

	err := json.Unmarshal(rawBytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil
}

func UnmarshalsDate(rawBytes []byte) ([]Date, error) {
	var data DateWrapper

	err := json.Unmarshal(rawBytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil
}

func UnmarshalsRelation(rawBytes []byte) ([]Relation, error) {
	var data RelationWrapper

	err := json.Unmarshal(rawBytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil
}


func LoadData() error{
	var err error
	 
	artistsRawBytes := FetchData("https://groupietrackers.herokuapp.com/api/artists")

	Artists, err = UnmarshalsArtist(artistsRawBytes)
	if err != nil {
		return err
	}
	for i := range Artists {
		name := strings.TrimSpace(Artists[i].Name)
		if strings.EqualFold(name, "Mamonas Assassinas") || strings.EqualFold(name, "Mamonas Assasinans") {
			Artists[i].Image = "/pict/mamonas.jpg"
		}
	}

	locationsRawBytes := FetchData("https://groupietrackers.herokuapp.com/api/locations")

	Locations, err = UnmarshalsLocation(locationsRawBytes)
	if err != nil {
		return err
	}

	datesRawBytes := FetchData("https://groupietrackers.herokuapp.com/api/dates")

	Dates, err = UnmarshalsDate(datesRawBytes)
	if err != nil {
		return err 
	}


	relationsRawBytes := FetchData("https://groupietrackers.herokuapp.com/api/relation")

	Relations, err = UnmarshalsRelation(relationsRawBytes)
	if err != nil {
		return err 
	}

	return nil
}
