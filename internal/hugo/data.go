package hugo

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type BookCollection map[string]Book

type Book struct {
	Title         string   `json:"title"`
	Series        string   `json:"series"`
	SeriesIndex   int      `json:"series_index"`
	Authors       []string `json:"authors"`
	PublishedYear int      `json:"published_year"`
	Publisher     string   `json:"publisher"`
	Isbn13        string   `json:"isbn_13"`
	Cover         string   `json:"cover"`
	Description   string   `json:"description"`
	Subjects      []string `json:"subjects"`
	Source        string   `json:"source"`
	ExternalIds   struct {
		Openlibrary string `json:"openlibrary"`
	} `json:"external_ids"`
}

func (h *Hugo) LoadBookData() BookCollection {
	jsonFilePath := path.Join(h.Config.dataDir, "books.json")
	file, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var bd BookCollection
	if err := json.Unmarshal(file, &bd); err != nil {
		log.Fatal(err)
	}
	return bd
}
