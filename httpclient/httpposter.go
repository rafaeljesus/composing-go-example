package httpclient

import (
	"io"
	"net/http"
)

type (
	// HTTPPoster holds fields and dependencies for executing an http POST request
	HTTPPoster struct {
		*Requester
	}
)

// NewHTTPPoster returns a configured HTTPPoster
func NewHTTPPoster(r *Requester) *HTTPPoster {
	return &HTTPPoster{r}
}

// Post executes a POST http request
func (c *HTTPPoster) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return c.Do(http.MethodPost, url, contentType, body)
}
