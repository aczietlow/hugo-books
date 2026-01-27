package openlibraryapi

import (
	"fmt"
	"strings"
)

type work struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Key         string `json:"key"`
	// "small": "https://covers.openlibrary.org/b/id/123-S.jpg",
	Covers  []int `json:"covers"`
	Authors []struct {
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
	} `json:"authors"`
	Subjects         []string `json:"subjects,omitempty"`
	FirstPublishDate string   `json:"first_publish_date"`
}

func (w *work) getWorksId() (string, error) {
	if strings.Contains(w.Key, "/works/") {
		return strings.Trim(w.Key, "/works/"), nil
	}
	return "", fmt.Errorf("%s not a valid key", w.Key)
}

// Finds and returns all open library author ids as part of a work
func (w *work) getAuthorIds() ([]string, error) {
	ids := []string{}

	for _, author := range w.Authors {
		if strings.Contains(author.Author.Key, "/authors/") {
			ids = append(ids, strings.Trim(author.Author.Key, "/authors/"))
		}
	}

	if len(ids) > 0 {
		return ids, nil
	}

	return []string{}, fmt.Errorf("couldn't find an author key on openlibrary work: %s", w.Key)
}
