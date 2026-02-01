package config

import (
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func Test_loadConfig(t *testing.T) {
	fsys := fstest.MapFS{
		"config.json": &fstest.MapFile{
			Data: []byte(`
hugo:
  basePath: "/var/www/blog"
  dataDir: "data"
  contentDir: "content"
  imageDir: "static/images/books"
openLibrary:
  httpTimeout: 10
  cacheTTL: 15
  userAgent: "HugoBooks/0.1 (aczietlow@gmail.com)"
  baseUrl: ""`),
		},
	}

	tests := []struct {
		testName string
		fsys     fs.FS
		name     string
		want     *Config
		wantErr  bool
	}{
		{
			testName: "Config loads correctly from file",
			fsys:     fsys,
			name:     "config.json",
			want: &Config{
				Hugo: hugoConfig{
					BasePath:   "/var/www/blog",
					DataDir:    "data",
					ContentDir: "content",
					ImageDir:   "static/images/books",
				},
				OpenLibrary: openLibraryConfig{
					HTTPTimeout: 10,
					CacheTTL:    15,
					UserAgent:   "HugoBooks/0.1 (aczietlow@gmail.com)",
					BaseUrl:     "",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf, gotErr := loadConfigFromFile(tt.fsys, tt.name)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("loadConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("loadConfig() succeeded unexpectedly")
			}
			if !reflect.DeepEqual(tt.want, conf) {
				t.Fatalf("config did not match expected!\nExpected: %v\nActual: %v\n", tt.want, conf)
			}
		})
	}
}
