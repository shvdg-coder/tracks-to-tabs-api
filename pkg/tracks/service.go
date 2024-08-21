package tracks

import (
	"github.com/google/uuid"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// Operations represent all operations related to tracks.
type Operations interface {
	DataOperations
	MappingOperations
	GetTracksCascading(trackID ...uuid.UUID) ([]*Track, error)
}

// Service is responsible for managing and retrieving tracks.
type Service struct {
	DataOperations
	MappingOperations
	TrackTabsOps  trktab.Operations
	TabsOps       tbs.Operations
	ReferencesOps ref.Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping MappingOperations, trackTabs trktab.Operations, tabs tbs.Operations, references ref.Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		TrackTabsOps:      trackTabs,
		TabsOps:           tabs,
		ReferencesOps:     references,
	}
}

// GetTracksCascading retrieves tabs, with entity references, for the provided IDs.
func (s Service) GetTracksCascading(trackID ...uuid.UUID) ([]*Track, error) {
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
