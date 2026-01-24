package openlibraryapi

import (
	"fmt"
	"strings"
)

// type author struct {
// 	Name string `json:"personal_name"`
// }

type openLibraryBook struct {
	Work     work
	Editions editions
}

type work struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Title       string   `json:"title"`
	Subjects    []string `json:"subjects"`
	Description string   `json:"description"`
	Key         string   `json:"key"`
	Covers      []int    `json:"covers"`
}

func (w *work) getWorksId() (string, error) {
	if strings.Contains(w.Key, "/books/") {
		return strings.Trim(w.Key, "/books/"), nil
	}
	return "", fmt.Errorf("%s not a valid key", w.Key)

}

type editions struct {
	Size    int       `json:"size"`
	Entries []edition `json:"entries"`
}

type edition struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	AuthorKeys []struct {
		Key string `json:"key"`
	} `json:"authors"`
	Authors   []string
	Languages []struct {
		Key string `json:"key"`
	} `json:"languages"`
	PublishDate string   `json:"publish_date"`
	Publishers  []string `json:"publishers"`
	Subjects    []string `json:"subjects,omitempty"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	FullTitle   string   `json:"full_title,omitempty"`
	Key         string   `json:"key"`
	Covers      []int    `json:"covers,omitempty"`
	Isbn13      []string `json:"isbn_13"`
}
