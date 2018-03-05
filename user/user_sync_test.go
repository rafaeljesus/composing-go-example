package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/rafaeljesus/composing-go-example/mock"
)

func TestUserSync(t *testing.T) {
	tests := []struct {
		scenario string
		function func(*testing.T, *mock.HTTPClientMock)
	}{
		{
			"test user syncer",
			testUserSync,
		},
		{
			"test fail to sync user",
			testFailToSyncUser,
		},
		{
			"test user syncer return error code",
			testUserSyncReturnErrorCode,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			c := new(mock.HTTPClientMock)
			test.function(t, c)
		})
	}
}

func testUserSync(t *testing.T, client *mock.HTTPClientMock) {
	client.PostFunc = func(url, contentType string, body io.Reader) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		if contentType == "" {
			t.Fatal("unexpected contentType")
		}
		u := new(User)
		if err := json.NewDecoder(body).Decode(u); err != nil {
			t.Fatalf("unexpected body decode error: %v", err)
		}
		if u.Country == "" {
			t.Fatal("unexpected country")
		}
		if u.Email == "" {
			t.Fatal("unexpected email")
		}
		return &http.Response{StatusCode: http.StatusOK}, nil
	}
	syncer := NewSync(client)
	u := New("foo@mail.com", "de")
	if err := syncer.Sync(u); err != nil {
		t.Fatalf("failed to sync user: %v", err)
	}
	if !client.PostInvoked {
		t.Fatal("expected client.Post() to be invoked")
	}
}

func testFailToSyncUser(t *testing.T, client *mock.HTTPClientMock) {
	client.PostFunc = func(url, contentType string, body io.Reader) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		if contentType == "" {
			t.Fatal("unexpected contentType")
		}
		u := new(User)
		if err := json.NewDecoder(body).Decode(u); err != nil {
			t.Fatalf("unexpected body decode error: %v", err)
		}
		if u.Country == "" {
			t.Fatal("unexpected country")
		}
		if u.Email == "" {
			t.Fatal("unexpected email")
		}
		return nil, errors.New("network error")
	}
	syncer := NewSync(client)
	u := New("foo@mail.com", "de")
	if err := syncer.Sync(u); err == nil {
		t.Fatal("expected client.Post() to return network error")
	}
	if !client.PostInvoked {
		t.Fatal("expected client.Post() to be invoked")
	}
}

func testUserSyncReturnErrorCode(t *testing.T, client *mock.HTTPClientMock) {
	client.PostFunc = func(url, contentType string, body io.Reader) (*http.Response, error) {
		if url == "" {
			t.Fatal("unexpected url")
		}
		if contentType == "" {
			t.Fatal("unexpected contentType")
		}
		u := new(User)
		if err := json.NewDecoder(body).Decode(u); err != nil {
			t.Fatalf("unexpected body decode error: %v", err)
		}
		if u.Country == "" {
			t.Fatal("unexpected country")
		}
		if u.Email == "" {
			t.Fatal("unexpected email")
		}
		return &http.Response{StatusCode: http.StatusInternalServerError}, nil
	}
	syncer := NewSync(client)
	u := New("foo@mail.com", "de")
	if err := syncer.Sync(u); err == nil {
		t.Fatal("expected client.Post() to return error code")
	}
	if !client.PostInvoked {
		t.Fatal("expected client.Post() to be invoked")
	}
}
