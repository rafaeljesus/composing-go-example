package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	syncURL = "http://localhost:8000/sync"
)

type (
	// Sync holds fields and dependencies for synchronizing users
	Sync struct {
		client HTTPPoster
	}
)

// NewSync returns a configured Sync
func NewSync(hp HTTPPoster) *Sync {
	return &Sync{hp}
}

// Sync responsible for synchronizing the user with third party system
func (s *Sync) Sync(user *User) error {
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(user); err != nil {
		return fmt.Errorf("failed to marshal user payload: %v", err)
	}

	res, err := s.client.Post(syncURL, "application/json", body)
	if err != nil {
		return fmt.Errorf("failed to sync user: %v", err)
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted:
		return nil
	default:
		return fmt.Errorf("failed to sync user: %d", res.StatusCode)
	}
}
