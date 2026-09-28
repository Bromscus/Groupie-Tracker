package functions

import (
	"strings"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"-"`
	Dates        []string            `json:"-"`
	Relations    map[string][]string `json:"-"`
}

type artistRecord struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type locationIndex struct {
	Index []Location `json:"index"`
}

type dateIndex struct {
	Index []Date `json:"index"`
}

type relationIndex struct {
	Index []Relation `json:"index"`
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

var Artists []Artist

func LoadData() error {
	artists, err := reader()
	if err != nil {
		return err
	}

	Artists = artists
	return nil
}

func reader() ([]Artist, error) {
	const baseURL = "https://groupietrackers.herokuapp.com/api"

	var api API
	if err := fetchJSON(baseURL, &api); err != nil {
		return nil, err
	}

	var records []artistRecord
	if err := fetchJSON(api.Artists, &records); err != nil {
		return nil, err
	}

	var locations locationIndex
	if err := fetchJSON(api.Locations, &locations); err != nil {
		return nil, err
	}

	var dates dateIndex
	if err := fetchJSON(api.Dates, &dates); err != nil {
		return nil, err
	}

	var relations relationIndex
	if err := fetchJSON(api.Relation, &relations); err != nil {
		return nil, err
	}

	locationsByID := make(map[int][]string, len(locations.Index))
	for _, location := range locations.Index {
		locationsByID[location.ID] = location.Locations
	}

	datesByID := make(map[int][]string, len(dates.Index))
	for _, date := range dates.Index {
		datesByID[date.ID] = date.Dates
	}

	relationsByID := make(map[int]map[string][]string, len(relations.Index))
	for _, relation := range relations.Index {
		relationsByID[relation.ID] = relation.DatesLocations
	}

	artists := make([]Artist, 0, len(records))

	for _, record := range records {
		artist := Artist{
			ID:           record.ID,
			Image:        record.Image,
			Name:         record.Name,
			Members:      record.Members,
			CreationDate: record.CreationDate,
			FirstAlbum:   record.FirstAlbum,
			Locations:    locationsByID[record.ID],
			Dates:        datesByID[record.ID],
			Relations:    relationsByID[record.ID],
		}

		if strings.EqualFold(artist.Name, "Mamonas Assassinas") {
			artist.Image = "style/Mamonas_Assassinas.jpg"
		}

		artists = append(artists, artist)
	}

	return artists, nil
}
