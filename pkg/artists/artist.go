package artists

import (
	"github.com/google/uuid"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// Artist represents an artist
type Artist struct {
	ID         uuid.UUID
	Name       string
	Tracks     []*trk.Track
	References []*ref.Reference
}

// Option modifies an Artist with configuration options.
type Option func(*Artist)

// WithID sets the ID of an Artist.
func WithID(id uuid.UUID) Option {
	return func(a *Artist) {
		a.ID = id
	}
}

// WithTracks sets the tracks of an Artist.
func WithTracks(tracks []*trk.Track) Option {
	return func(a *Artist) {
		a.Tracks = tracks
	}
}

// NewArtist instantiates a new Artist.
func NewArtist(name string, configs ...Option) *Artist {
	artist := &Artist{ID: uuid.New(), Name: name}
	for _, config := range configs {
		config(artist)
	}
	return artist
}
