package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {

	books := scanHugoContentForISBN()
	for _, book := range books {
		fmt.Printf("%v\n", book)
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
