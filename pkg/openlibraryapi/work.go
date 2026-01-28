package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetWork(id string, httpClient *http.Client) (work, error) {
	id = strings.ToUpper(id)

	idType := getIdType(id)

	w := work{}
	var err error

	switch idType {

	case int(openLibraryWork):
		w, err = getWorkByOlId(id, httpClient)
	}

	if err != nil {
		return work{}, err
	}

	return w, nil
}

// Fetches open library work by open library work id.
// works/{{ openLibraryId }}
func getWorkByOlId(id string, httpClient *http.Client) (work, error) {
	url := baseUrl + "/works/" + id + ".json"

	resp, err := httpClient.Get(url)
	if err != nil {
		return work{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return work{}, fmt.Errorf("received a %d reponse from the api\n", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return work{}, err
	}

	w := work{}
	if err := json.Unmarshal(body, &w); err != nil {
		return work{}, err
	}
	return w, nil
}
