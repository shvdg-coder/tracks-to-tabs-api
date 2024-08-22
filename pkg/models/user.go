package models

import "github.com/google/uuid"

// UserEntry represents a user.
type UserEntry struct {
	ID    uuid.UUID
	Email string
}
