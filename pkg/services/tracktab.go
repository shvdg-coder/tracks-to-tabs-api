package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/schemas"
)

// TrackTabOps represents all operations related to 'track to tab' links.
type TrackTabOps interface {
	schemas.TrackTabSchema
	data.TrackTabData
	CreateTrackTabEntries(track *models.TrackEntry, tabs []*models.TabEntry) []*models.TrackTabEntry
	ExtractIDsFromTrackTabEntries(trackTabs []*models.TrackTabEntry) (trackIDs []uuid.UUID, tabIDs []uuid.UUID)
}

// TrackTabSvc is responsible for managing and retrieving 'track to tab' links.
type TrackTabSvc struct {
	schemas.TrackTabSchema
	data.TrackTabData
}

// NewTrackTabSvc instantiates a TrackTabSvc.
func NewTrackTabSvc(schema schemas.TrackTabSchema, data data.TrackTabData) TrackTabOps {
	return &TrackTabSvc{
		TrackTabSchema: schema,
		TrackTabData:   data}
}

// CreateTrackTabEntries creates models.TrackTab's using the provided models.TrackEntry and models.TabEntry's.
func (tt *TrackTabSvc) CreateTrackTabEntries(track *models.TrackEntry, tabs []*models.TabEntry) []*models.TrackTabEntry {
	trackTabs := make([]*models.TrackTabEntry, len(tabs))
	for i, tab := range tabs {
		trackTabs[i] = &models.TrackTabEntry{TrackID: track.ID, TabID: tab.ID}
	}
	return trackTabs
}

// ExtractIDsFromTrackTabEntries retrieves the track and tab IDs from the models.TrackTabEntry's.
func (tt *TrackTabSvc) ExtractIDsFromTrackTabEntries(trackTabs []*models.TrackTabEntry) (trackIDs []uuid.UUID, tabIDs []uuid.UUID) {
	trackIDs = make([]uuid.UUID, len(trackTabs))
	tabIDs = make([]uuid.UUID, len(trackTabs))

	for _, trackTab := range trackTabs {
		trackIDs = append(trackIDs, trackTab.TrackID)
		tabIDs = append(tabIDs, trackTab.TabID)
	}

	return trackIDs, tabIDs
}
