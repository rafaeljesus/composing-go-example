package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	usersURL = "http://localhost:8081/users"
)

type (
	// Store holds fields and dependencies for storing users
	Store struct {
		client HTTPClient
	}
)

// NewStore returns a configured Store
func NewStore(hc HTTPClient) *Store {
	return &Store{hc}
}

// Store responsible for storing the user
func (s *Store) Store(user *User) error {
	userResource := fmt.Sprintf("%s/%s/%s", usersURL, user.Email, user.Country)
	res, err := s.client.Get(userResource)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %v", err)
	}

	if res.StatusCode == http.StatusNotFound {
		return errors.New("user not found")
	}

	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(user); err != nil {
		return fmt.Errorf("failed to marshal user payload: %v", err)
	}

	res, err = s.client.Post(usersURL, "application/json", body)
	if err != nil {
		return fmt.Errorf("failed to store user: %v", err)
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	default:
		return errors.New("failed to store user")
	}
}
