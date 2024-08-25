package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// TrackTabOps represents all operations related to 'track to tab' links.
type TrackTabOps interface {
	schemas.TrackTabSchema
	data.TrackTabData
	ExtractIDsFromTrackTabEntries(trackTabs []*models.TrackTabEntry) ([]uuid.UUID, []uuid.UUID)
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

// ExtractIDsFromTrackTabEntries retrieves the track and tab IDs from the models.TrackTabEntry's.
func (t *TrackTabSvc) ExtractIDsFromTrackTabEntries(trackTabs []*models.TrackTabEntry) ([]uuid.UUID, []uuid.UUID) {
	var trackIDs []uuid.UUID
	var tabIDs []uuid.UUID
	for _, trackTab := range trackTabs {
		trackIDs = append(trackIDs, trackTab.TrackID)
		tabIDs = append(tabIDs, trackTab.TabID)
	}
	return trackIDs, tabIDs
}
