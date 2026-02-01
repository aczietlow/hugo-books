package main

import (
	"fmt"
	"log"

	"github.com/aczietlow/hugo-books/internal/config"
	"github.com/aczietlow/hugo-books/internal/hugo"
	"github.com/aczietlow/hugo-books/pkg/openlibraryapi"
)

func main() {
	conf, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	hugo := hugo.NewHugo(conf)
	bookAPI := openlibraryapi.NewClient(conf)
	// TODO: Use existing collection data to seed bookcache
	collection := hugo.LoadBookData()

	for _, book := range hugo.ScanHugoContentForBooks() {
		if _, exists := collection[book.ISBN]; !exists {
			fmt.Printf("Fetching details for %s using isbn:%v\n", book.Title, book.ISBN)
			b, err := bookAPI.GetBookById(book.ISBN)
			if err != nil {
				log.Fatal(err)
			}
			collection[book.ISBN] = b
		}
	}

	err = hugo.SaveBookData(collection)
	if err != nil {
		log.Fatal(err)
	}

	//https://covers.openlibrary.org/b/id/14627060.jpg
	bookAPI.FetchCoverById(14627060)
}
