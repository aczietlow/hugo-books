package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aczietlow/hugo-books/internal/hugo"
	"github.com/aczietlow/hugo-books/pkg/openlibraryapi"
)

func main() {
	hugo := hugo.NewHugo("/home/aczietlow/Projects/hugo-blog/")
	bookAPI := openlibraryapi.NewClient(10*time.Second, 15*time.Minute, "")
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

	err := hugo.SaveBookData(collection)
	if err != nil {
		log.Fatal(err)
	}

}
