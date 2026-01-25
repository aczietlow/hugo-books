package openlibraryapi

import "testing"

func Test_getIdType(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want int
	}{
		{
			name: "OpenLibrary Works ID",
			id:   "OL27513W",
			want: int(openLibraryWork),
		},
		{
			name: "OpenLibrary Edition ID",
			id:   "OL26452600M",
			want: int(openLibraryEdition),
		},
		{
			name: "OpenLibrary Author",
			id:   "OL26452600A",
			want: int(openLibraryAuthor),
		},
		{
			name: "ISBN 10",
			id:   "0547928211",
			want: int(isbn10),
		},
		{
			name: "ISBN 13",
			id:   "9780547928210",
			want: int(isbn13),
		},
		{
			name: "ISBN 13",
			id:   "978054-7928210",
			want: int(isbn13),
		},
		{
			name: "unknown",
			id:   "fellowshipofring0000tolk_o5y1",
			want: int(unknown),
		},
		{
			name: "case and space insensitivity",
			id:   " ol27513w   ",
			want: int(openLibraryWork),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIdType(tt.id)
			if got != tt.want {
				t.Errorf("getIdType() = %v, want %v", got, tt.want)
			}
		})
	}
}
