package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TrackTabOps represents all operations related to 'track to tab' links.
type TrackTabOps interface {
	data.TrackTabData
	ExtractTabIDs(trackTabs []*models.TrackTabEntry) []uuid.UUID
}

// TrackTabSvc is responsible for managing and retrieving 'track to tab' links.
type TrackTabSvc struct {
	data.TrackTabData
}

// NewTrackTabSvc instantiates a TrackTabSvc.
func NewTrackTabSvc(data data.TrackTabData) TrackTabOps {
	return &TrackTabSvc{TrackTabData: data}
}

// ExtractTabIDs retrieves the tab IDs from the models.TrackTabEntry's.
func (t *TrackTabSvc) ExtractTabIDs(trackTabs []*models.TrackTabEntry) []uuid.UUID {
	var tabIDs []uuid.UUID
	for _, link := range trackTabs {
		tabIDs = append(tabIDs, link.TabID)
	}
	return tabIDs
}
