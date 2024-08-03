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
	trktab.Operations
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

// LinkTrackToTab creates a link between a track and a tab using their IDs.
func (s Service) LinkTrackToTab(trackID, tabID uuid.UUID) {
	s.TrackTabsOps.LinkTrackToTab(trackID, tabID)
}

// GetTrackToTabLink retrieves the tab linked to a track using the track's ID.
func (s Service) GetTrackToTabLink(trackID uuid.UUID) (*trktab.TrackTab, error) {
	return s.TrackTabsOps.GetTrackToTabLink(trackID)
}

// GetTrackToTabLinks retrieves the links between a list of track IDs and their associated tabs.
func (s Service) GetTrackToTabLinks(trackID ...uuid.UUID) ([]*trktab.TrackTab, error) {
	return s.TrackTabsOps.GetTrackToTabLinks(trackID...)
}

// ExtractTabIDs retrieves the tab IDs from each tracktab.TrackTab.
func (s Service) ExtractTabIDs(trackTabs []*trktab.TrackTab) []uuid.UUID {
	return s.TrackTabsOps.ExtractTabIDs(trackTabs)
}

// GetTracksCascading retrieves tabs, with entity references, for the provided IDs.
func (s Service) GetTracksCascading(tabID ...uuid.UUID) ([]*Track, error) {
	tracks, err := s.GetTracks(tabID...)
	if err != nil {
		return nil, err
	}
	trackTabs, err := s.GetTrackToTabLinks(tabID...)
	if err != nil {
		return nil, err
	}
	tabIDs := s.ExtractTabIDs(trackTabs)
	tabs, err := s.TabsOps.GetTabs(tabIDs...)
	if err != nil {
		return nil, err
	}
	tracksMap := s.ToMap(tracks)
	tabsMap := s.TabsOps.ToMap(tabs)
	tracks = s.MapTabsToTracks(trackTabs, tracksMap, tabsMap)
	return tracks, nil
}
