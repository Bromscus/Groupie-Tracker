package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchJSON(url string, info any) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: %d from %s", response.StatusCode, url)
	}

	return json.NewDecoder(response.Body).Decode(info)
}
