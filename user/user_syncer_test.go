package user

import (
	"io"
	"net/http"
	"testing"

	"github.com/rafaeljesus/composing-go-example/mock"
)

func TestUserSync(t *testing.T) {
	client := new(mock.HTTPPosterMock)
	client.PostFunc = func(url, contentType string, body io.Reader) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		if contentType == "" {
			t.Fatal("unexpected contentType")
		}
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
	syncer := NewUserSyncer(client)
	u := New("foo@mail.com", "de")
	if err := syncer.Sync(u); err != nil {
		t.Fatalf("failed to sync user: %v", err)
	}
	if !client.PostInvoked {
		t.Fatal("expected client.Post() to be invoked")
	}
}
