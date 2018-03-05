package mock

import (
	"io"
	"net/http"
)

type (
	HTTPClientMock struct {
		HTTPPostMock
		HTTPGetMock
	}

	HTTPPostMock struct {
		PostInvoked bool
		PostFunc    func(url, contentType string, body io.Reader) (*http.Response, error)
	}

	HTTPGetMock struct {
		GetInvoked bool
		GetFunc    func(url string) (*http.Response, error)
	}
)

func (m *HTTPPostMock) Post(url, contentType string, body io.Reader) (*http.Response, error) {
	m.PostInvoked = true
	return m.PostFunc(url, contentType, body)
}

func (m *HTTPGetMock) Get(url string) (*http.Response, error) {
	m.GetInvoked = true
	return m.GetFunc(url)
}
