package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	syncURL = "http://localhost:8000/sync"
)

type (
	// UserSyncer holds fields and dependencies for synchronizing users
	UserSyncer struct {
		client HTTPPoster
	}
)

// NewUserSyncer returns a configured UserSyncer
func NewUserSyncer(hp HTTPPoster) *UserSyncer {
	return &UserSyncer{hp}
}

// Sync responsible for synchronizing the user with third party system
func (s *UserSyncer) Sync(user *User) error {
	body, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user payload: %v", err)
	}

	res, err := s.client.Post(syncURL, contentType, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to sync user: %v", err)
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	default:
		return errors.New("failed to sync user")
	}
}
