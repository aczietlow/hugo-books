package openlibraryapi

import (
	"net/http"
	"time"

	"github.com/aczietlow/hugo-books/internal/config"
	"github.com/aczietlow/hugo-books/pkg/bookcache"
)

const baseUrl = "https://openlibrary.org"

type Client struct {
	httpClient http.Client
	baseUrl    string
	cache      bookcache.Cache
}

type Transport struct {
	UserAgent string
	Transport http.RoundTripper
}

// Creates a new client to be reused for all api requests
//
// Sets identifiable UserAgent as requested by openlibrary api docs
func NewClient(c *config.Config) Client {
	httpTimeout := time.Duration(c.OpenLibrary.HTTPTimeout) * time.Second
	cacheTTL := time.Duration(c.OpenLibrary.CacheTTL) * time.Minute
	return Client{
		httpClient: http.Client{
			Transport: &Transport{
				UserAgent: c.OpenLibrary.UserAgent,
				Transport: http.DefaultTransport,
			},
			Timeout: httpTimeout,
		},
		baseUrl: getBaseUrl(c.OpenLibrary.BaseUrl),
		cache:   bookcache.NewCacheStorage(cacheTTL),
	}
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	if t.UserAgent != "" {
		r.Header.Set("User-Agent", t.UserAgent)
	}
	return t.Transport.RoundTrip(r)
}

func getBaseUrl(url string) string {
	if url != "" {
		return url
	}
	return baseUrl
}
