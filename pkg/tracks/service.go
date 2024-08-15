package tracks

import (
	"github.com/google/uuid"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// Operations represent all operations related to tracks.
type Operations interface {
	DataOperations
	MappingOperations
	GetTracksCascading(tabID ...uuid.UUID) ([]*Track, error)
}

// Service is responsible for managing and retrieving tracks.
type Service struct {
	DataOperations
	MappingOperations
	TrackTabsOps trktab.Operations
	TabsOps      tbs.Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping MappingOperations, trackTabs trktab.Operations, tabs tbs.Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		TrackTabsOps:      trackTabs,
		TabsOps:           tabs,
	}
}

// GetTracksCascading retrieves tabs, with entity references, for the provided IDs.
func (s Service) GetTracksCascading(tabID ...uuid.UUID) ([]*Track, error) {
	tracks, err := s.GetTracks(tabID...)
	if err != nil {
		return nil, err
	}
	trackTabs, err := s.TrackTabsOps.GetTrackToTabLinks(tabID...)
	if err != nil {
		return nil, err
	}
	tabIDs := s.TrackTabsOps.ExtractTabIDs(trackTabs)
	tabs, err := s.TabsOps.GetTabs(tabIDs...)
	if err != nil {
		return nil, err
	}
	tracksMap := s.TracksToMap(tracks)
	tabsMap := s.TabsOps.TabsToMap(tabs)
	tracks = s.MapTabsToTracks(trackTabs, tracksMap, tabsMap)
	return tracks, nil
}
