package main

import (
	"fmt"
	"log"
	"path"

	"github.com/aczietlow/hugo-books/internal/config"
	"github.com/aczietlow/hugo-books/internal/hugo"
	"github.com/aczietlow/hugo-books/pkg/openlibraryapi"
)

func main() {
	fetchImages := true
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	hugo := hugo.NewHugo(conf)
	bookAPI := openlibraryapi.NewClient(conf)
	// TODO: Use existing collection data to seed bookcache
	collection := hugo.LoadBookData()

	for _, book := range hugo.ScanHugoContentForBooks() {
		if _, exists := collection[book.Isbn]; !exists {
			fmt.Printf("Fetching details for %s using isbn:%v\n", book.Title, book.Isbn)
			b, err := bookAPI.GetBookById(book.Isbn)
			if err != nil {
				log.Fatal(err)
			}
			if fetchImages && b.CoverId != "" {
				imageDir := path.Join(conf.Hugo.BasePath, conf.Hugo.ImageDir)
				err := bookAPI.FetchCoverById(b.CoverId, imageDir)
				if err != nil {
					log.Fatal(err)
				}

			}

			collection[book.Isbn] = b
		}

		// // Temp trying to figure out file fetching...
		// b := collection[book.ISBN]
		// imageDir := path.Join(conf.Hugo.BasePath, conf.Hugo.ImageDir)
		// imageFile := path.Join(imageDir, b.CoverId+".jpg")
		//
		// _, err = os.Stat(imageFile)
		//
		// if err != nil && !errors.Is(err, fs.ErrExist) {
		// 	fmt.Printf("Missing cover image for %s, attempting to fetch: %s\n", b.Title, imageFile)
		// 	err := bookAPI.FetchCoverById(b.CoverId, imageDir)
		// 	if err != nil {
		// 		fmt.Printf("ecountered error  during fetch")
		// 		log.Fatal(err)
		// 	}
		// }
	}

	err = hugo.SaveBookData(collection)
	if err != nil {
		log.Fatal(err)
	}
}
