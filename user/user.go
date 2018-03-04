package user

import (
	"io"
	"net/http"
)

var (
	contentType = "application/json"
)

type (
	// User represents the user
	User struct {
		Email   string `json:"email"`
		Country string `json:"country"`
	}

	// HTTPClient is the http wrapper for the application
	HTTPClient interface {
		HTTPGetter
		HTTPPoster
	}

	// HTTPPoster holds fields and dependencies for executing an http POST request
	HTTPPoster interface {
		// Post executes a POST http request
		Post(url, contentType string, body io.Reader) (*http.Response, error)
	}

	// HTTPGetter holds fields and dependencies for executing an http GET request
	HTTPGetter interface {
		// Get executes a GET http request
		Get(url string) (*http.Response, error)
	}
)

// New returns new user
func New(email, country string) *User {
	return &User{
		Email:   email,
		Country: country,
	}
}
