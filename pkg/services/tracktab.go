package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents all operations related to 'track to tab' links.
type Operations interface {
	database.TabsOps
	ExtractTabIDs(trackTabs []*models.TrackTabEntry) []uuid.UUID
}

// ArtistTrackSvc is responsible for managing and retrieving 'track to tab' links.
type Service struct {
	database.TabsOps
}

// NewTrackSvc instantiates a ArtistTrackSvc.
func NewService(data database.TabsOps) Operations {
	return &Service{TabsOps: data}
}

// ExtractTabIDs retrieves the tab IDs from each TrackTabEntry.
func (service *Service) ExtractTabIDs(trackTabs []*models.TrackTabEntry) []uuid.UUID {
	var tabIDs []uuid.UUID
	for _, link := range trackTabs {
		tabIDs = append(tabIDs, link.TabID)
	}
	return tabIDs
}
