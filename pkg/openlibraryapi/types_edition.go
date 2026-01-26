package openlibraryapi

import (
	"fmt"
	"strings"
)

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
	Works []struct {
		Key string `json:"key"`
	} `json:"works"`
	Languages []struct {
		Key string `json:"key"`
	} `json:"languages"`
	Authors     []string
	PublishDate string   `json:"publish_date"`
	Publishers  []string `json:"publishers"`
	Subjects    []string `json:"subjects,omitempty"`
	Title       string   `json:"title"`
	Series      []string `json:"series,omitempty"`
	Subtitle    string   `json:"subtitle"`
	FullTitle   string   `json:"full_title,omitempty"`
	Key         string   `json:"key"`
	Covers      []int    `json:"covers,omitempty"`
	Isbn13      []string `json:"isbn_13"`
	Isbn10      []string `json:"isbn_10"`
}

func (e *edition) getWorksId() (string, error) {
	if len(e.Works) > 0 {
		if strings.Contains(e.Works[0].Key, "/works/") {
			return strings.Trim(e.Works[0].Key, "/works/"), nil
		}
	}
	return "", fmt.Errorf("Could not find a valid work id")
}
