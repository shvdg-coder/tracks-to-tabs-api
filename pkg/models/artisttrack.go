package models

import "github.com/google/uuid"

// ArtistTrackEntry represents a 'artist to track' link from the database.
type ArtistTrackEntry struct {
	ArtistID uuid.UUID
	TrackID  uuid.UUID
}

// Fields returns a slice of interfaces containing values of the ArtistTrackEntry.
func (at *ArtistTrackEntry) Fields() []interface{} {
	return []interface{}{at.ArtistID, at.TrackID}
}

// ArtistTrack represents an 'artist to track' link with entity references.
type ArtistTrack struct {
	Artist *Artist
	Track  *Track
}
