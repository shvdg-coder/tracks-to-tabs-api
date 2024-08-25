package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)
import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// ArtistTrackOps combines the interface representing operations related to 'artist to track' links.
type ArtistTrackOps interface {
	schemas.ArtistTrackSchema
	data.ArtistTrackData
	ExtractIDsFromArtistTrackEntries(artistTracks []*models.ArtistTrackEntry) (artistIDs []uuid.UUID, trackIDs []uuid.UUID)
}

// ArtistTrackSvc is responsible for managing 'artist to track' links.
type ArtistTrackSvc struct {
	schemas.ArtistTrackSchema
	data.ArtistTrackData
}

// NewArtistTrackSvc instantiates an ArtistTrackSvc.
func NewArtistTrackSvc(schema schemas.ArtistTrackSchema, data data.ArtistTrackData) ArtistTrackOps {
	return &ArtistTrackSvc{
		ArtistTrackSchema: schema,
		ArtistTrackData:   data,
	}
}

// ExtractIDsFromArtistTrackEntries retrieves the artist and track IDs from the models.ArtistTrackEntry's.
func (at *ArtistTrackSvc) ExtractIDsFromArtistTrackEntries(artistTracks []*models.ArtistTrackEntry) (artistIDs []uuid.UUID, trackIDs []uuid.UUID) {
	artistIDs = make([]uuid.UUID, len(artistTracks))
	trackIDs = make([]uuid.UUID, len(artistTracks))

	for _, artistTrack := range artistTracks {
		artistIDs = append(artistIDs, artistTrack.ArtistID)
		trackIDs = append(trackIDs, artistTrack.TrackID)
	}

	return artistIDs, trackIDs
}
