package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	collection := loadBookData()
	for _, details := range collection {
		fmt.Println(details.Title)
	}

	books := scanHugoContentForISBN()
	for _, b := range books {
		if _, exist := collection[b.ISBN]; !exist {
			fmt.Printf("Need to fetch data for %s\n", b.Title)
		}
	}

}

type BookContent struct {
	Path  string
	Title string
	ISBN  string
}

func scanHugoContentForISBN() []BookContent {
	fileSystem := os.DirFS("/home/aczietlow/Projects/hugo-blog/content/books")

	books := []BookContent{}
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path == "." || path == ".." {
			return nil
		}

		file, err := fs.ReadFile(fileSystem, path)
		if err != nil {
			log.Fatal(err)
		}
		bm, err := getFrontMatter(file)
		if err != nil {
			log.Fatal(err)
		}

		if bm.ISBN == "" {
			return nil
		} else {
			bm.Path = path
		}
		books = append(books, bm)

		return nil
	})

	return books
}

// Stupidest parser... in the world
func getFrontMatter(file []byte) (BookContent, error) {
	bm := BookContent{}
	content := strings.Split(string(file), "\n")
	fmMarker := 0
	for _, l := range content {
		if l == "+++" {
			fmMarker++
			continue
		}

		if fmMarker == 1 {
			if strings.HasPrefix(l, "isbn") {
				start := strings.Index(l, "\"")
				last := strings.LastIndex(l, "\"")
				bm.ISBN = l[start+1 : last]
			}
			if strings.HasPrefix(l, "title") {
				start := strings.Index(l, "\"")
				last := strings.LastIndex(l, "\"")
				bm.Title = l[start+1 : last]
			}
		}
	}

	return bm, nil
}

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

func loadBookData() BookCollection {
	file, err := os.ReadFile("/home/aczietlow/Projects/hugo-blog/data/books.json")
	if err != nil {
		log.Fatal(err)
	}

	var bd BookCollection
	if err := json.Unmarshal(file, &bd); err != nil {
		log.Fatal(err)
	}
	return bd
}
