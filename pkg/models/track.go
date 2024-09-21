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

// Fields returns a slice of interfaces containing values of the TrackEntry.
func (t *TrackEntry) Fields() []interface{} {
	return []interface{}{t.ID.String(), t.Title, t.Duration}
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
