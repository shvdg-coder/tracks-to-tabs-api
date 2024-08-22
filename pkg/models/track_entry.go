package models

import (
	"github.com/google/uuid"
)

// TrackEntry represents a track in the database.
type TrackEntry struct {
	ID       uuid.UUID
	Title    string
	Duration uint // In milliseconds.
}
