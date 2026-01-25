package openlibraryapi

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		httpTimeout time.Duration
		cacheTTL    time.Duration
		baseUrl     string
		// want        openlibraryapi.Client
		userAgent string
	}{
		{
			name:        "User Agent Should be set on every request",
			httpTimeout: 100 * time.Millisecond,
			cacheTTL:    100 * time.Millisecond,
			baseUrl:     "http://localhost:11001",
			userAgent:   "HugoBooks/0.1 (aczietlow@gmail.com)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.httpTimeout, tt.cacheTTL, tt.baseUrl)

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

			if string(body) != tt.userAgent {
				t.Errorf("Unexpected User-Agent String! Found: %s - Expected %s", string(body), tt.userAgent)
			}
		})
	}
}
