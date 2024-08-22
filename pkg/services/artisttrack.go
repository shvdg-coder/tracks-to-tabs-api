package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents all operations related to 'artist to track' links.
type Operations interface {
	database.TrackOps
	ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID
}

// ArtistTrackSvc is responsible for managing and retrieving 'artist to track' links.
type Service struct {
	database.TrackOps
}

// NewTrackSvc instantiates a ArtistTrackSvc.
func NewService(data database.TrackOps) Operations {
	return &Service{TrackOps: data}
}

// ExtractTrackIDs retrieves the track IDs from each ArtistTrackEntry.
func (s *Service) ExtractTrackIDs(artistTracks []*models.ArtistTrackEntry) []uuid.UUID {
	var trackIDs []uuid.UUID
	for _, link := range artistTracks {
		trackIDs = append(trackIDs, link.TrackID)
	}
	return trackIDs
}
