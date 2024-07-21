package artisttrack

import "github.com/google/uuid"

// ArtistTrack represents a link from an artist to a track.
type ArtistTrack struct {
	ArtistID uuid.UUID
	TrackID  uuid.UUID
}
