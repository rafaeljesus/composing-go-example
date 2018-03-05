package httpclient

import (
	"io"
	"net/http"
)

type (
	// HTTPClient is the http wrapper for the application
	HTTPClient struct {
		*Request
	}
)

// New returns a configured HTTPClient
func New(r *Request) *HTTPClient {
	return &HTTPClient{r}
}

// Get executes a GET http request
func (c *HTTPClient) Get(url string) (*http.Response, error) {
	return c.Do(http.MethodGet, url, "application/json", nil)
}

// Post executes a POST http request
func (c *HTTPClient) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Do(http.MethodPost, url, contentType, body)
}
