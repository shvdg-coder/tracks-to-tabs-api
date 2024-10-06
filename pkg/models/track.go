package models

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
)

// TrackEntry represents a track in the database.
type TrackEntry struct {
	ID       uuid.UUID      `db:"id"`
	Title    string         `db:"title"`
	Duration uint           `db:"duration"` // In milliseconds.
	Cover    sql.NullString `db:"cover"`    // Image
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
