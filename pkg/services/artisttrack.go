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
	ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID
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

// ExtractTrackIDs retrieves the track IDs from the models.ArtistTrackEntry's.
func (at *ArtistTrackSvc) ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID {
	var trackIDs []uuid.UUID
	for _, link := range artistTracks {
		trackIDs = append(trackIDs, link.TrackID)
	}
	return trackIDs
}
