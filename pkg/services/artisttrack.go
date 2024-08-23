package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents all operations related to 'artist to track' links.
type ArtistTrackOps interface {
	data.ArtistTrackData
	ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID
}

// ArtistTrackSvc is responsible for managing and retrieving 'artist to track' links.
type ArtistTrackSvc struct {
	data.ArtistTrackData
}

// NewArtistTrackSvc instantiates a ArtistTrackSvc.
func NewArtistTrackSvc(data data.ArtistTrackData) ArtistTrackOps {
	return &ArtistTrackSvc{ArtistTrackData: data}
}

// ExtractTrackIDs retrieves the track IDs from the models.ArtistTrackEntry's.
func (at *ArtistTrackSvc) ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID {
	var trackIDs []uuid.UUID
	for _, link := range artistTracks {
		trackIDs = append(trackIDs, link.TrackID)
	}
	return trackIDs
}
