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
	Series []string `json:"series,omitempty"`
	Works  []struct {
		Key string `json:"key"`
	} `json:"works"`
	// Do I need lang on editions?
	Languages []struct {
		Key string `json:"key"`
	} `json:"languages"`
	PublishDate string   `json:"publish_date"`
	Publishers  []string `json:"publishers"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Key         string   `json:"key"`
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

func (e *edition) getEditionId() (string, error) {
	if strings.Contains(e.Key, "/books/") {
		return strings.Trim(e.Key, "/books/"), nil
	}
	return "", fmt.Errorf("Could not find a valid edition id")
}
