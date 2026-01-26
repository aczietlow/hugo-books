package openlibraryapi

// type author struct {
// 	Name string `json:"personal_name"`
// }

type openLibraryBook struct {
	Work     work
	Editions editions
	Edition  edition
}

// Book object we want to ship back to the client
type book struct {
	Title         string
	Subtitle      string
	Authors       []string
	Description   string
	ISBN          string
	Genre         []string
	Cover         string
	Source        string
	Series        string
	SeriesIndex   int
	PublishedYear int
	Publisher     string
	externalIds   map[string]string
}
