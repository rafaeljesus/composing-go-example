package httpclient

import "net/http"

type (
	// HTTPGetter holds fields and dependencies for executing an http GET request
	HTTPGetter struct {
		*Requester
	}
)

// NewHTTPPoster returns a configured HTTPPoster
func NewHTTPGetter(r *Requester) *HTTPGetter {
	return &HTTPGetter{r}
}

// Get executes a GET http request
func (c *HTTPGetter) Get(url string) (*http.Response, error) {
	return c.Do(http.MethodGet, url, contentType, nil)
}
