package models

import "github.com/google/uuid"

// UserEntry represents a user.
type UserEntry struct {
	ID    uuid.UUID
	Email string
}

// User represents a user with entity references.
type User struct {
	*UserEntry
}
