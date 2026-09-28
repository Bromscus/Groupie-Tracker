package functions

import "testing"

func TestArtistMatchesSearch(t *testing.T) {
	artist := Artist{
		Name:       "Queen",
		Members:    []string{"Freddie Mercury", "Brian May"},
		Locations:  []string{"London"},
		Dates:      []string{"1970-01-01"},
		FirstAlbum: "Queen",
	}

	tests := []struct {
		name  string
		query string
		want  bool
	}{
		{name: "empty query", query: "", want: true},
		{name: "matches name", query: "queen", want: true},
		{name: "matches member", query: "freddie", want: true},
		{name: "matches location", query: "london", want: true},
		{name: "no match", query: "notfound", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := artistMatchesSearch(artist, tt.query); got != tt.want {
				t.Fatalf("artistMatchesSearch(%q) = %v, want %v", tt.query, got, tt.want)
			}
		})
	}
}
