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
	// Syncer holds fields and dependencies for synchronizing users
	Syncer struct {
		client HTTPPoster
	}
)

// NewSyncer returns a configured Syncer
func NewSyncer(hp HTTPPoster) *Syncer {
	return &Syncer{hp}
}

// Sync responsible for synchronizing the user with third party system
func (s *Syncer) Sync(user *User) error {
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
