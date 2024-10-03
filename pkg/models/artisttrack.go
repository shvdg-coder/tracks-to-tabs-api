package models

import "github.com/google/uuid"

// ArtistTrackEntry represents a 'artist to track' link from the database.
type ArtistTrackEntry struct {
	ArtistID uuid.UUID `db:"artist_id"`
	TrackID  uuid.UUID `db:"track_id"`
}

// ArtistTrack represents an 'artist to track' link with entity references.
type ArtistTrack struct {
	Artist *Artist
	Track  *Track
}
