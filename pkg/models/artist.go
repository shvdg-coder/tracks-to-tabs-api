package models

import "github.com/google/uuid"

// ArtistEntry represents an artist in the database.
type ArtistEntry struct {
	ID   uuid.UUID
	Name string
}

// Artist represents an artist with entity references.
type Artist struct {
	*ArtistEntry
	Tracks     []*Track
	References []*Reference
}
