package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getEdition(id string, httpClient *http.Client) (edition, error) {
	id = strings.ToUpper(id)

	idType := getIdType(id)

	e := edition{}
	var err error

	switch idType {

	case int(openLibraryEdition):
		e, err = getEditionByOlId(id, httpClient)
	case int(isbn10):
	case int(isbn13):
		e, err = getEditionByIsbn(id, httpClient)
	default:
		err = fmt.Errorf("Illegal id type used when fetching edition. id: %s, type: %d", id, idType)
	}

	if err != nil {
		return edition{}, err
	}

	return e, nil
}

// books/{{ openLibraryId }}.json
func getEditionByOlId(id string, httpClient *http.Client) (edition, error) {
	url := baseUrl + "/books/" + id + ".json"

	resp, err := httpClient.Get(url)
	if err != nil {
		return edition{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return edition{}, fmt.Errorf("received a %d reponse from the api\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return edition{}, err
	}

	e := edition{}
	if err := json.Unmarshal(body, &e); err != nil {
		return edition{}, err
	}
	return e, nil
}

// isbn/{{ isbn10||isbn13 }}.json
func getEditionByIsbn(id string, httpClient *http.Client) (edition, error) {
	url := baseUrl + "/isbn/" + id + ".json"
	resp, err := httpClient.Get(url)
	if err != nil {
		return edition{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return edition{}, fmt.Errorf("Received a %d reponse from the api\n GET %s\n", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return edition{}, err
	}

	e := edition{}
	if err := json.Unmarshal(body, &e); err != nil {
		return edition{}, err
	}
	return e, nil
}

// Gets all available editions for a work.
// Uses openlibrary work id
func getWorkEditions(id string, httpClient *http.Client) (editions, error) {
	url := baseUrl + "/works/" + id + "/editions.json"
	resp, err := httpClient.Get(url)
	if err != nil {
		return editions{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return editions{}, fmt.Errorf("received a %d reponse from the api\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return editions{}, err
	}

	e := editions{}
	if err := json.Unmarshal(body, &e); err != nil {
		return editions{}, err
	}

	// Only return english editions
	// TODO: Reslice existing slice to reduce memory footprint
	e2 := editions{
		Size:    0,
		Entries: []edition{},
	}

	// for _, edition := range e.Entries {
	// 	if len(edition.Languages) > 0 && edition.Languages[0].Key == "/languages/eng" {
	// 		for _, author := range edition.AuthorKeys {
	// 			a, err := getAuthorByKey(author.Key, httpClient)
	// 			if err != nil {
	// 				return editions{}, err
	// 			}
	// 			edition.Authors = append(edition.Authors, a.Name)
	//
	// 		}
	// 		e2.Entries = append(e2.Entries, edition)
	// 		e2.Size++
	// 	}
	// }

	return e2, nil
}
