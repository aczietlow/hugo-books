package openlibraryapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aczietlow/hugo-books/internal/config"
)

func TestNewClient(t *testing.T) {
	testConf, _ := config.LoadConfigFromBytes([]byte(`
hugo:
  basePath: /test
  dataDir: data
  contentDir: content
  imageDir: images
openLibrary:
  httpTimeout: 1
  cacheTTL: 0
  userAgent: test
  baseUrl: http://fake
`))
	tests := []struct {
		name      string
		conf      *config.Config
		userAgent string
	}{
		{
			name: "User Agent Should be set on every request",
			conf: testConf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client := NewClient(tt.conf)

			userAgentHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				io.WriteString(w, req.Header.Get("user-agent"))
			})

			s := httptest.NewServer(userAgentHandler)
			defer s.Close()

			resp, err := client.httpClient.Get(s.URL)
			if err != nil {
				t.Logf("Error when client made request to server: %v\n", err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Logf("Received a %d reponse from the api\n", resp.StatusCode)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Logf("Failed to read body: %v", err)
			}

			if string(body) != tt.conf.OpenLibrary.UserAgent {
				t.Errorf("Unexpected User-Agent String! Found: %s - Expected %s", string(body), tt.conf.OpenLibrary.UserAgent)
			}
		})
	}
}
