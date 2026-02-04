package hugo

import (
	"reflect"
	"testing"
)

func Test_getFrontMatter(t *testing.T) {
	testContent := `+++
title = "Digital Minimalism" 
date = 2025-01-16T19:37:34-05:00
startDate = 2025-01-16
endDate = 2025-02-03
isbn = "9780525536512"
+++

### Recap

Digital Minimalism, by Cal Newport, focused on applying the ideals of the minimalism moment to a chaotic, digital world. I picked this up as a continuation of my interest and focus on achieving a happier life. Cal would use the phrase "deep life." Digital minimalism isn't about technological luddite, so much as it is about spending our limited time on higher areas of return. This really further highlights that our time on this Earth and specifically our time is a zero sum game.  
`
	malFormatted := `+++
title = "Digital Minimalism" 
isbn = "978-052-553-6512"
+++
`

	tests := []struct {
		name    string
		file    []byte
		expect  contentFrontMatter
		wantErr bool
	}{
		{
			name: "Can extract title from Front matter",
			file: []byte(testContent),
			expect: contentFrontMatter{
				Title: "Digital Minimalism",
				Isbn:  "9780525536512",
			},
			wantErr: false,
		},
		{
			name: "Correctly format isbn number",
			file: []byte(malFormatted),
			expect: contentFrontMatter{
				Title: "Digital Minimalism",
				Isbn:  "9780525536512",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm, gotErr := getFrontMatter(tt.file)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getFrontMatter() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getFrontMatter() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(fm, tt.expect) {
				t.Errorf("getFrontMatter() = %v, expect %v", fm, tt.expect)
			}
		})
	}
}
