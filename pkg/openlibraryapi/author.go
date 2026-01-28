package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Fetches Author by openlibrary author id
func getAuthorById(id string, httpClient *http.Client) (author, error) {
	url := baseUrl + "/authors/" + id + ".json"
	resp, err := httpClient.Get(url)
	if err != nil {
		return author{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return author{}, fmt.Errorf("received a %d reponse from the api\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return author{}, err
	}

	a := author{}
	if err := json.Unmarshal(body, &a); err != nil {
		return author{}, err
	}

	return a, nil
}
