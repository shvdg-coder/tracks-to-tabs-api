package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// TrackEntry represents a track in the database.
type TrackEntry struct {
	ID       uuid.UUID
	Title    string
	Duration uint // In milliseconds.
}

// Track represents a track.
type Track struct {
	*TrackEntry
	Artist     *Artist
	Tabs       []*Tab
	References []*Reference
}

// MarshalJSON marshals the models.Track while preventing circling.
func (t *Track) MarshalJSON() ([]byte, error) {
	track := *t
	track.Artist = &Artist{
		ArtistEntry: t.Artist.ArtistEntry,
	}
	return json.Marshal(track)
}
