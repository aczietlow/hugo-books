package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// TODO: what is this struct for?
type book struct {
	Title       string
	Subtitle    string
	Authors     []string
	Description string
	ISBN        string
	Genre       []string
	Cover       string
	Source      string
}

// Fetches a book by OpenLibraryID or ISBN number
func (c *Client) GetBookById(id string) (book, error) {
	id = strings.ToUpper(id)

	// Attempt to fetch from cache first
	if cacheRecord, exists := c.cache.Get(id); exists {
		lr := openLibraryBook{}
		if err := json.Unmarshal(cacheRecord, &lr); err != nil {
			return book{}, nil
		}
		b := aggregateLibraryRecord(lr)
		return b, nil
	}

	lr, err := getBookDetails(id, &c.httpClient)
	if err != nil {
		return book{}, err
	}

	cacheRecord, err := json.Marshal(lr)
	if err != nil {
		return book{}, err
	}
	c.cache.Add(id, cacheRecord)

	b := aggregateLibraryRecord(lr)
	return b, nil
}

func aggregateLibraryRecord(libraryRecord openLibraryBook) book {
	b := book{
		Title: libraryRecord.Work.Title,
	}

	// Set description if available
	if libraryRecord.Work.Description != "" {
		b.Description = libraryRecord.Work.Description
	}

	if libraryRecord.Work.Key != "" {
		b.Source = baseUrl + libraryRecord.Work.Key
	}

	mappedFields := []string{}

	// Loop through each edition looking for the data to populate a whole book object.
	for _, edition := range libraryRecord.Editions.Entries {
		if b.Subtitle == "" && edition.Subtitle != "" {
			b.Subtitle = edition.Subtitle
			mappedFields = append(mappedFields, "Subtitle")
		}

		if b.ISBN == "" && len(edition.Isbn13) > 0 && edition.Isbn13[0] != "" {
			// Assume we'll only ever want a single ISBN number
			b.ISBN = edition.Isbn13[0]
			mappedFields = append(mappedFields, "ISBN")
		}

		if len(b.Genre) <= 0 && len(edition.Subjects) > 0 {
			b.Genre = edition.Subjects
			mappedFields = append(mappedFields, "Genre")
		}

		if b.Cover == "" && len(edition.Covers) > 0 {
			b.Cover = "https://covers.openlibrary.org/b/id/" + strconv.Itoa(edition.Covers[0]) + ".jpg"
			mappedFields = append(mappedFields, "Cover")
		}

		if len(b.Authors) <= 0 && len(edition.Authors) > 0 {
			b.Authors = edition.Authors
			mappedFields = append(mappedFields, "Authors")
		}

		// Stop iterating throgh editions if we have all the data required.
		if len(mappedFields) == 5 {
			break
		}
	}

	return b
}

func getBookDetails(id string, httpClient *http.Client) (openLibraryBook, error) {
	libraryRecord := openLibraryBook{}

	e, err := getEdition(id, httpClient)
	if err != nil {
		return openLibraryBook{}, err
	}

	libraryRecord.Edition = e

	olWorksId, err := e.getWorksId()
	if err != nil {
		return openLibraryBook{}, err
	}

	w, err := getWorkByOlId(olWorksId, httpClient)
	if err != nil {
		return openLibraryBook{}, err
	}

	libraryRecord.Work = w

	allEditions, err := getWorkEditions(olWorksId, httpClient)
	if err != nil {
		return openLibraryBook{}, err
	}
	libraryRecord.Editions = allEditions

	return libraryRecord, nil
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

	for _, edition := range e.Entries {
		if len(edition.Languages) > 0 && edition.Languages[0].Key == "/languages/eng" {
			for _, author := range edition.AuthorKeys {
				a, err := getAuthorByKey(author.Key, httpClient)
				if err != nil {
					return editions{}, err
				}
				edition.Authors = append(edition.Authors, a.Name)

			}
			e2.Entries = append(e2.Entries, edition)
			e2.Size++
		}
	}

	return e2, nil
}

func getAuthorByKey(key string, httpClient *http.Client) (author, error) {
	url := baseUrl + "/" + key + ".json"
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
		return author{}, nil
	}

	return a, nil
}
