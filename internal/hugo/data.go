package hugo

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/aczietlow/hugo-books/pkg/openlibraryapi"
)

type BookCollection map[string]openlibraryapi.Book

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

func (h *Hugo) SaveBookData(collection BookCollection) error {
	jsonFilePath := path.Join(h.Config.dataDir, "books.json")
	encodedData, err := json.MarshalIndent(collection, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFilePath, encodedData, 0666)
	if err != nil {
		return err
	}
	return nil
}
