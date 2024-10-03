package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// TrackEntry represents a track in the database.
type TrackEntry struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"title"`
	Cover    string    `db:"cover"`    // Image
	Duration uint      `db:"duration"` // In milliseconds.
}

// Track represents a track.
type Track struct {
	*TrackEntry
	Artist     *Artist
	Tabs       []*Tab
	References []*Reference
	Resources  []*Resource
}

// MarshalJSON marshals the models.Track while preventing circling.
func (t *Track) MarshalJSON() ([]byte, error) {
	track := *t
	track.Artist = &Artist{
		ArtistEntry: t.Artist.ArtistEntry,
		References:  t.Artist.References,
		Resources:   t.Artist.Resources,
		Tracks:      nil,
	}
	return json.Marshal(track)
}
