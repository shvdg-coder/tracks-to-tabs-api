package artists

import (
	"github.com/google/uuid"
)

// Artist represents an artist
type Artist struct {
	ID   uuid.UUID
	Name string
}

// NewArtist instantiates a new Artist.
func NewArtist(name string) *Artist {
	return &Artist{
		ID:   uuid.New(),
		Name: name,
	}
}
