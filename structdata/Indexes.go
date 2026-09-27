package structdata

import "strings"

var (
	detailsByID    map[int]ArtistDetails
	lowercaseNames []string
)

func buildIndexes() {
	locationsByID := make(map[int]Location, len(Locations))
	for _, location := range Locations {
		if _, exists := locationsByID[location.ID]; !exists {
			locationsByID[location.ID] = location
		}
	}
	datesByID := make(map[int]Date, len(Dates))
	for _, date := range Dates {
		if _, exists := datesByID[date.ID]; !exists {
			datesByID[date.ID] = date
		}
	}
	relationsByID := make(map[int]Relation, len(Relations))
	for _, relation := range Relations {
		if _, exists := relationsByID[relation.ID]; !exists {
			relationsByID[relation.ID] = relation
		}
	}

	detailsByID = make(map[int]ArtistDetails, len(Artists))
	lowercaseNames = make([]string, len(Artists))
	for i, artist := range Artists {
		lowercaseNames[i] = strings.ToLower(artist.Name)
		location, hasLocation := locationsByID[artist.ID]
		date, hasDate := datesByID[artist.ID]
		relation, hasRelation := relationsByID[artist.ID]
		if _, exists := detailsByID[artist.ID]; !exists && hasLocation && hasDate && hasRelation {
			detailsByID[artist.ID] = ArtistDetails{Artist: artist, Location: location, Date: date, Relation: relation}
		}
	}
}

func GetArtistDetails(id int) (ArtistDetails, bool) {
	detail, found := detailsByID[id]
	return detail, found
}

func SearchArtists(query string) []Artist {
	matched := make([]Artist, 0)
	for i, name := range lowercaseNames {
		if strings.Contains(name, query) {
			matched = append(matched, Artists[i])
		}
	}
	return matched
}
