package tracks

import (
	"github.com/google/uuid"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// Operations represent all operations related to tracks.
type Operations interface {
	DatabaseOperations
	MappingOperations
	trktab.Operations
}

// Service is responsible for managing and retrieving tracks.
type Service struct {
	DatabaseOperations
	MappingOperations
	TrackTabsOps trktab.Operations
}

// NewService instantiates a Service.
func NewService(database DatabaseOperations, mapping MappingOperations, trackTabs trktab.Operations) Operations {
	return &Service{
		DatabaseOperations: database,
		MappingOperations:  mapping,
		TrackTabsOps:       trackTabs,
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
