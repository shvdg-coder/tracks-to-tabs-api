package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// TrackOps represent all operations related to tracks.
type TrackOps interface {
	schemas.TrackSchema
	data.TrackData
	mappers.TrackMapper
	GetTracks(trackID ...uuid.UUID) ([]*models.Track, error)
}

// TrackSvc is responsible for managing and retrieving tracks.
type TrackSvc struct {
	schemas.TrackSchema
	data.TrackData
	mappers.TrackMapper
	TrackTabOps
	TabOps
	ReferenceOps
}

// NewTrackSvc instantiates a TrackSvc.
func NewTrackSvc(schema schemas.TrackSchema, data data.TrackData, mapper mappers.TrackMapper, trackTabs TrackTabOps, tabs TabOps, references ReferenceOps) TrackOps {
	return &TrackSvc{
		TrackSchema:  schema,
		TrackData:    data,
		TrackMapper:  mapper,
		TrackTabOps:  trackTabs,
		TabOps:       tabs,
		ReferenceOps: references,
	}
}

// GetTracks retrieves tabs, with entity references, for the provided IDs.
func (t TrackSvc) GetTracks(trackID ...uuid.UUID) ([]*models.Track, error) {
	trackEntries, err := t.GetTrackEntries(trackID...)
	if err != nil {
		return nil, err
	}

	trackTabEntries, err := t.GetTrackToTabLinks(trackID...)
	if err != nil {
		return nil, err
	}

	tabIDs := t.ExtractTabIDs(trackTabEntries)
	tabs, err := t.GetTabs(tabIDs...)
	if err != nil {
		return nil, err
	}

	references, err := t.GetReferences(trackID...)
	if err != nil {
		return nil, err
	}

	tracks := t.TrackEntriesToTracks(trackEntries)
	tracksMap := t.TracksToMap(tracks)
	tabsMap := t.TabsToMap(tabs)

	tracksMap = t.MapTabsToTracks(tracksMap, tabsMap, trackTabEntries)
	tracksMap = t.MapReferencesToTracks(tracksMap, references)
	tracks = t.MapToTracks(tracksMap)

	return tracks, nil
}
