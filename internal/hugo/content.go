package hugo

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

type BookContent struct {
	Path  string
	Title string
	ISBN  string
}

// Walks throgh Hugo Content directory. searching for book conent
//
// Book Content is determined to be anything that has "ISBN" property in the front matter.
//
// Returns a slice of Book Hugo Content
func (h *Hugo) ScanHugoContentForBooks() []BookContent {
	fileSystem := os.DirFS(h.Config.contentDir)

	books := []BookContent{}
	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if path == "." || path == ".." || d.IsDir() == true {
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
