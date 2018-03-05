package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

type (
	// Request is the application http request
	// let's use this and not repeat our self in every implementation
	Request struct{}
)

// Do executes http request
func (r *Request) Do(method, url, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request %v: ", err)
	}

	req.Header.Set("Content-Type", contentType)

	return http.DefaultClient.Do(req)
}
