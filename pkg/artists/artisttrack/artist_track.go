package artisttrack

import "github.com/google/uuid"

// ArtistTrack represents a link from an artist to a track.
type ArtistTrack struct {
	artistID uuid.UUID
	trackID  uuid.UUID
}
