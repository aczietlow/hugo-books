package openlibraryapi

import (
	"fmt"
	"strings"
)

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
