package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TrackOps represent all operations related to tracks.
type TrackOps interface {
	data.TrackData
	mappers.TrackMapper
	GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error)
}

// TrackSvc is responsible for managing and retrieving tracks.
type TrackSvc struct {
	data.TrackData
	mappers.TrackMapper
	TrackTabOps
	TabOps
	ReferenceOps
}

// NewTrackSvc instantiates a TrackSvc.
func NewTrackSvc(data data.TrackData, mapper mappers.TrackMapper, trackTabs TrackTabOps, tabs TabOps, references ReferenceOps) TrackOps {
	return &TrackSvc{
		TrackData:    data,
		TrackMapper:  mapper,
		TrackTabOps:  trackTabs,
		TabOps:       tabs,
		ReferenceOps: references,
	}
}

// GetTracksCascading retrieves tabs, with entity references, for the provided IDs.
func (t TrackSvc) GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error) {
	return nil, nil
}
