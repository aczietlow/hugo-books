package openlibraryapi

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type idType int

const (
	unknown idType = iota
	openLibraryWork
	openLibraryEdition
	openLibraryAuthor
	isbn10
	isbn13
)

func PrettyPrint(data any) {
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("failed to pretty print map")
	}
	fmt.Println(string(encodedData))
}

func getIdType(id string) int {
	var i int

	id = strings.ToUpper(strings.TrimSpace(id))

	// Openlibrary Identifiers
	if strings.HasPrefix(id, "OL") {
		if strings.HasSuffix(id, "W") {
			return int(openLibraryWork)
		} else if strings.HasSuffix(id, "M") {
			return int(openLibraryEdition)
		} else if strings.HasSuffix(id, "A") {
			return int(openLibraryAuthor)
		}
	}

	// remove any dashes that might be present in isbn numbers
	id = strings.ReplaceAll(id, "-", "")
	isNumber := func(c rune) bool {
		return c < '0' || c > '9'
	}

	// ISBN identifiers
	if strings.IndexFunc(id, isNumber) == -1 {
		if len(id) == 10 {
			return int(isbn10)
		} else if len(id) == 13 {
			return int(isbn13)
		}
	}

	return i
}
