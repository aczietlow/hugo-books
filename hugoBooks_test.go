package main

import "testing"

func Test_getFrontMatter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		file    []byte
		want    BookContent
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := getFrontMatter(tt.file)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getFrontMatter() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getFrontMatter() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("getFrontMatter() = %v, want %v", got, tt.want)
			}
		})
	}
}
