package mock

import (
	"io"
	"net/http"
)

type (
	HTTPClientMock struct {
		HTTPPosterMock
		HTTPGetterMock
	}

	HTTPPosterMock struct {
		PostInvoked bool
		PostFunc    func(url, contentType string, body io.Reader) (*http.Response, error)
	}

	HTTPGetterMock struct {
		GetInvoked bool
		GetFunc    func(url string) (*http.Response, error)
	}
)

func (m *HTTPPosterMock) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	m.PostInvoked = true
	return m.PostFunc(url, contentType, body)
}

func (m *HTTPGetterMock) Get(url string) (*http.Response, error) {
	m.GetInvoked = true
	return m.GetFunc(url)
}
