package tracks

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// MappingOperations represents operations related to data mapping.
type MappingOperations interface {
	GetTracksCascading(artistID ...string) ([]*Track, error)
}

// MappingService is responsible for mapping entities to tracks.
type MappingService struct {
	DatabaseOperations
	TracksTabsOps trktab.Operations
	TracksOps     tabs.Operations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService(tracks DatabaseOperations, trackTabs trktab.Operations, tabs tabs.Operations) MappingOperations {
	return &MappingService{
		DatabaseOperations: tracks,
		TracksTabsOps:      trackTabs,
		TracksOps:          tabs}
}

// GetTracksCascading retrieves tracks, with entity references, for the provided IDs.
func (m *MappingService) GetTracksCascading(trackID ...string) ([]*Track, error) {
	return nil, nil
}
