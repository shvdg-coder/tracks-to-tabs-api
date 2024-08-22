package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TrackOps represent all operations related to tracks.
type Operations interface {
	DataOperations
	mappers.MappingOperations
	GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error)
}

// Service is responsible for managing and retrieving tracks.
type Service struct {
	DataOperations
	mappers.MappingOperations
	TrackTabsOps  Operations
	TabsOps       Operations
	ReferencesOps Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping mappers.MappingOperations, trackTabs Operations, tabs Operations, references Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		TrackTabsOps:      trackTabs,
		TabsOps:           tabs,
		ReferencesOps:     references,
	}
}

// GetTracksCascading retrieves tabs, with entity references, for the provided IDs.
func (s Service) GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error) {
	tracks, err := s.GetTracks(trackID...)
	if err != nil {
		return nil, err
	}

	trackTabs, err := s.TrackTabsOps.GetTrackToTabLinks(trackID...)
	if err != nil {
		return nil, err
	}

	tabIDs := s.TrackTabsOps.ExtractTabIDs(trackTabs)
	tabs, err := s.TabsOps.GetTabsCascading(tabIDs...)
	if err != nil {
		return nil, err
	}

	references, err := s.ReferencesOps.GetReferencesCascading(trackID...)
	if err != nil {
		return nil, err
	}

	tracksMap := s.TracksToMap(tracks)
	tabsMap := s.TabsOps.TabsToMap(tabs)
	tracks = s.MapTabsToTracks(trackTabs, tracksMap, tabsMap)
	tracks = s.MapReferencesToTracks(tracksMap, references)

	return tracks, nil
}
