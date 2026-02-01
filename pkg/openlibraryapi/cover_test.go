package openlibraryapi

import "testing"

func Test_buildCoverImageUrl(t *testing.T) {
	tests := []struct {
		name    string
		coverId int
		want    string
	}{
		{
			name:    "build correct cover image url",
			coverId: 14627060,
			want:    "https://covers.openlibrary.org/b/id/14627060.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildCoverImageUrl(tt.coverId)
			if got != tt.want {
				t.Errorf("buildCoverImageUrl() = expect:%v\nwant:%v\n", got, tt.want)
			}
		})
	}
}
