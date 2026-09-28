package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	functions "groupie-tracker/handlers"
)

func TestSearchHandlerReturnsArtists(t *testing.T) {
	projectRoot, err := filepath.Abs("..")
	if err != nil {
		t.Fatalf("failed to resolve project root: %v", err)
	}

	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	if err := os.Chdir(projectRoot); err != nil {
		t.Fatalf("failed to change to project root: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(oldWd)
	})

	functions.Artists = []functions.Artist{
		{ID: 1, Name: "Queen", Members: []string{"Freddie Mercury", "Brian May"}, CreationDate: 1970, FirstAlbum: "Queen"},
		{ID: 2, Name: "The Beatles", Members: []string{"John Lennon"}, CreationDate: 1960, FirstAlbum: "Please Please Me"},
	}

	req := httptest.NewRequest(http.MethodGet, "/search?q=queen", nil)
	w := httptest.NewRecorder()

	functions.SearchHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("SearchHandler() status = %d, want %d", w.Code, http.StatusOK)
	}

	body := w.Body.String()
	if body == "" {
		t.Fatal("SearchHandler() returned an empty response body")
	}
	if !strings.Contains(body, "Queen") {
		t.Fatalf("SearchHandler() response body does not contain expected artist name: %s", body)
	}
}
