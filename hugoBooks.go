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
	collection := hugo.LoadBookData()
	for _, details := range collection {
		fmt.Println(details.Title)
	}

	books := hugo.ScanHugoContentForBooks()
	for _, b := range books {
		if _, exist := collection[b.ISBN]; !exist {
			fmt.Printf("Need to fetch data for %s\n", b.Title)
		}
	}

	bookAPI := openlibraryapi.NewClient(6*time.Second, 15*time.Minute, "")
	// fmt.Printf("fetching %s\n", oid)
	b, err := bookAPI.GetBookById("9780547928210")
	if err != nil {
		log.Fatal(err)
	}

	openlibraryapi.PrettyPrint(b)
}
