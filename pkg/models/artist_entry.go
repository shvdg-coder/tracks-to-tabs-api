package models

import (
	"github.com/google/uuid"
)

// ArtistEntry represents an artist in the database.
type ArtistEntry struct {
	ID   uuid.UUID
	Name string
}
