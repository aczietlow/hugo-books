package openlibraryapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Fetches a book by OpenLibraryID or ISBN number
func (c *Client) GetBookById(id string) (Book, error) {
	id = strings.ToUpper(id)

	// Attempt to fetch from cache first
	if cacheRecord, exists := c.cache.Get(id); exists {
		lr := openLibraryBook{}
		if err := json.Unmarshal(cacheRecord, &lr); err != nil {
			return Book{}, nil
		}
		b := aggregateLibraryRecord(lr)
		return b, nil
	}

	lr, err := getBookDetails(id, &c.httpClient)
	if err != nil {
		return Book{}, err
	}

	cacheRecord, err := json.Marshal(lr)
	if err != nil {
		return Book{}, err
	}
	c.cache.Add(id, cacheRecord)

	b := aggregateLibraryRecord(lr)
	return b, nil
}

func aggregateLibraryRecord(libraryRecord openLibraryBook) Book {
	b := Book{
		Title:  libraryRecord.Work.Title,
		Source: "openlibrary",
	}

	// Set description if available
	if libraryRecord.Work.Description.Value != "" {
		b.Description = libraryRecord.Work.Description.Value
	}

	if libraryRecord.Edition.Subtitle != "" {
		b.Subtitle = libraryRecord.Edition.Subtitle
	}

	if len(libraryRecord.Edition.Isbn10) > 0 && libraryRecord.Edition.Isbn10[0] != "" {
		b.ISBN10 = libraryRecord.Edition.Isbn10[0]
	}

	if len(libraryRecord.Edition.Isbn13) > 0 && libraryRecord.Edition.Isbn13[0] != "" {
		b.ISBN13 = libraryRecord.Edition.Isbn13[0]
	}

	if len(libraryRecord.Work.Subjects) > 0 {
		// Limit genre to 5 at most
		limit := min(len(libraryRecord.Work.Subjects), 5)
		b.Genre = libraryRecord.Work.Subjects[0:limit]
	}

	if len(libraryRecord.Edition.Series) > 0 {
		b.Series = libraryRecord.Edition.Series[0]
	}

	if libraryRecord.Work.FirstPublishDate != "" {
		b.PublishedDate = libraryRecord.Work.FirstPublishDate
	}

	if libraryRecord.Edition.PublishDate != "" {
		b.PublishedDate = libraryRecord.Edition.PublishDate
	}

	if len(libraryRecord.Edition.Publishers) > 0 {
		b.Publishers = strings.Join(libraryRecord.Edition.Publishers, ",")
	}

	for _, author := range libraryRecord.Authors {
		b.Authors = append(b.Authors, author.Name)
	}

	wId, err := libraryRecord.Work.getWorksId()
	if err != nil {
		log.Fatalf("Failed fetching workd id: %v\n", err)
	}

	b.externalIds = map[string]string{
		"openlibraryWork": wId,
	}

	if len(libraryRecord.Work.Covers) > 0 {
		b.CoverUrl = buildCoverImageUrl(libraryRecord.Work.Covers[0])
		b.CoverId = strconv.Itoa(libraryRecord.Work.Covers[0])
	}

	// // Loop through each edition looking for the data to populate a whole book object.
	// mappedFields := []string{}
	// for _, edition := range libraryRecord.Editions.Entries {
	// 	if b.Subtitle == "" && edition.Subtitle != "" {
	// 		b.Subtitle = edition.Subtitle
	// 		mappedFields = append(mappedFields, "Subtitle")
	// 	}
	//
	// 	if b.ISBN == "" && len(edition.Isbn13) > 0 && edition.Isbn13[0] != "" {
	// 		// Assume we'll only ever want a single ISBN number
	// 		b.ISBN = edition.Isbn13[0]
	// 		mappedFields = append(mappedFields, "ISBN")
	// 	}
	//
	// 	if len(b.Genre) <= 0 && len(edition.Subjects) > 0 {
	// 		b.Genre = edition.Subjects
	// 		mappedFields = append(mappedFields, "Genre")
	// 	}
	//
	// 	if b.Cover == "" && len(edition.Covers) > 0 {
	// 		b.Cover = "https://covers.openlibrary.org/b/id/" + strconv.Itoa(edition.Covers[0]) + ".jpg"
	// 		mappedFields = append(mappedFields, "Cover")
	// 	}
	//
	// 	if len(b.Authors) <= 0 && len(edition.Authors) > 0 {
	// 		b.Authors = edition.Authors
	// 		mappedFields = append(mappedFields, "Authors")
	// 	}
	//
	// 	// Stop iterating through editions if we have all the data required.
	// 	if len(mappedFields) == 5 {
	// 		break
	// 	}
	// }

	return b
}

func getBookDetails(id string, httpClient *http.Client) (openLibraryBook, error) {
	libraryRecord := openLibraryBook{}

	e, err := getEdition(id, httpClient)
	if err != nil {
		return openLibraryBook{}, err
	}

	libraryRecord.Edition = e

	olWorksId, err := e.getWorksId()
	if err != nil {
		return openLibraryBook{}, err
	}

	w, err := getWorkByOlId(olWorksId, httpClient)
	if err != nil {
		log.Printf("attempted to fetch id: %s", olWorksId)
		return openLibraryBook{}, err
	}

	libraryRecord.Work = w

	allEditions, err := getWorkEditions(olWorksId, httpClient)
	if err != nil {
		return openLibraryBook{}, err
	}
	libraryRecord.Editions = allEditions

	olAuthorIds, err := w.getAuthorIds()
	if err != nil {
		return openLibraryBook{}, err
	}

	for _, authorId := range olAuthorIds {
		author, err := getAuthorById(authorId, httpClient)
		if err != nil {
			return openLibraryBook{}, err
		}
		libraryRecord.Authors = append(libraryRecord.Authors, author)
	}

	return libraryRecord, nil
}
