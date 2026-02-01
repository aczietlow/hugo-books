package openlibraryapi

// type author struct {
// 	Name string `json:"personal_name"`
// }

type openLibraryBook struct {
	Work     work
	Editions editions
	Edition  edition
	Authors  []author
}

// Book object we want to ship back to the client
type Book struct {
	Title         string
	Subtitle      string
	Series        string
	Authors       []string
	Description   string
	PublishedDate string
	Publishers    string
	CoverId       int
	CoverUrl      string
	Genre         []string
	ISBN10        string
	ISBN13        string
	Source        string
	// SeriesIndex   int
	externalIds map[string]string
}
