package models

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
)

// ArtistEntry represents an artist in the database.
type ArtistEntry struct {
	ID     uuid.UUID      `db:"id"`
	Name   string         `db:"name"`
	Cover  sql.NullString `db:"cover"`  // Image
	Banner sql.NullString `db:"banner"` // Image
}

// Artist represents an artist with entity references.
type Artist struct {
	*ArtistEntry
	Tracks     []*Track
	References []*Reference
	Resources  []*Resource
}

// MarshalJSON marshals the models.Artist.
func (a *Artist) MarshalJSON() ([]byte, error) {
	return json.Marshal(*a)
}
