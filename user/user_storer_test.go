package user

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/rafaeljesus/composing-go-example/mock"
)

func TestUserStore(t *testing.T) {
	client := new(mock.HTTPClientMock)
	client.PostFunc = func(url, contentType string, body io.Reader) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		if contentType == "" {
			t.Fatal("unexpected contentType")
		}
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
	client.GetFunc = func(url string) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: nopCloser{
				bytes.NewBufferString(`{"email": "foo@mail.com", "country": "de"}`),
			},
		}, nil
	}
	storer := NewUserStorer(client)
	u := New("foo@mail.com", "de")
	if err := storer.Store(u); err != nil {
		t.Fatalf("failed to store user: %v", err)
	}
	if !client.PostInvoked {
		t.Fatal("expected client.Post() to be invoked")
	}
	if !client.GetInvoked {
		t.Fatal("expected client.Get() to be invoked")
	}
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
