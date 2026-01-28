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
type book struct {
	Title         string
	Subtitle      string
	Authors       []string
	Description   string
	ISBN10        string
	ISBN13        string
	Genre         []string
	CoverUrl      string
	Source        string
	Series        string
	SeriesIndex   int
	PublishedDate string
	Publishers    string
	externalIds   map[string]string
}
