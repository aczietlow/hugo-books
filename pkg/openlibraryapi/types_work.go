package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type work struct {
	Type struct {
		Key string `json:"key"`
	} `json:"type"`
	Title       string           `json:"title"`
	Description descriptionField `json:"description"`
	Key         string           `json:"key"`
	Covers      []int            `json:"covers"`
	Authors     []struct {
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
	} `json:"authors"`
	Subjects         []string `json:"subjects,omitempty"`
	FirstPublishDate string   `json:"first_publish_date"`
}

type description struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type descriptionField struct {
	Description *description
	Value       string
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

// work description fields can either be string OR struct{type: string, value: string}
func (d *descriptionField) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err == nil {
		d.Value = value
		d.Description = nil
		return nil
	}
	var desc description
	if err := json.Unmarshal(data, &desc); err == nil {
		d.Value = desc.Value
		d.Description = &desc
		return nil
	}
	return fmt.Errorf("description field must be either type string or {type: string, value: string}, got %s", string(data))
}
