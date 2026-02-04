package hugo

import (
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type BookContent struct {
	Path  string
	Title string
	Isbn  string
}

// Walks throgh Hugo Content directory. searching for book conent
//
// Book Content is determined to be anything that has "ISBN" property in the front matter.
//
// Returns a slice of Book Hugo Content
func (h *Hugo) ScanHugoContentForBooks() []BookContent {
	path := path.Join(h.Config.hugoPath, h.Config.contentDir)
	fileSystem := os.DirFS(path)

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
		fm, err := getFrontMatter(file)
		if err != nil {
			log.Fatal(err)
		}

		if fm.Isbn == "" {
			return nil
		}
		bm := BookContent{
			Path:  path,
			Title: fm.Title,
			Isbn:  string(fm.Isbn),
		}

		books = append(books, bm)

		return nil
	})

	return books
}

type contentFrontMatter struct {
	Title string
	Isbn  Isbn
}

type Isbn string

// Stupidest parser... in the world
func getFrontMatter(file []byte) (contentFrontMatter, error) {
	content := strings.Split(string(file), "\n")

	// TODO: This currently assumes front matter will Always be toml
	frontMatter := ""
	fmMarker := 0
	for _, l := range content {
		if strings.Contains(l, "+++") {
			fmMarker++
			continue
		}

		if fmMarker == 1 {
			frontMatter += l + "\n"
		}
	}

	var fm contentFrontMatter

	err := toml.Unmarshal([]byte(frontMatter), &fm)
	if err != nil {
		return contentFrontMatter{}, err
	}

	return fm, nil
}

func (i *Isbn) UnmarshalText(data []byte) error {
	*i = Isbn(strings.TrimSpace(strings.ReplaceAll(string(data), "-", "")))
	return nil
}
