package tracks

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// MappingOperations represents operations related to track data mapping.
type MappingOperations interface {
	GetTracksAsMap(trackID ...uuid.UUID) (map[uuid.UUID]*Track, error)
	GetTracksCascading(trackID ...uuid.UUID) ([]*Track, error)
	GetTracksCascadingAsMap(trackID ...uuid.UUID) (map[uuid.UUID]*Track, error)
}

// MappingService is responsible for mapping entities to tracks.
type MappingService struct {
	DatabaseOperations
	TracksTabsOps trktab.Operations
	TabsOps       tabs.Operations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService(tracks DatabaseOperations, trackTabs trktab.Operations, tabs tabs.Operations) MappingOperations {
	return &MappingService{
		DatabaseOperations: tracks,
		TracksTabsOps:      trackTabs,
		TabsOps:            tabs}
}

// GetTracksAsMap retrieves the tracks for the provided IDs, and creates a map where the key is the ID and the value the Track.
func (m *MappingService) GetTracksAsMap(trackID ...uuid.UUID) (map[uuid.UUID]*Track, error) {
	tracks, err := m.GetTracks(trackID...)
	if err != nil {
		return nil, err
	}

	trackMap := make(map[uuid.UUID]*Track)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}

	return trackMap, nil
}

// GetTracksCascading retrieves tracks, with entity references, for the provided IDs.
func (m *MappingService) GetTracksCascading(trackID ...uuid.UUID) ([]*Track, error) {
	// Get tracks
	tracksMap, err := m.GetTracksAsMap(trackID...)
	if err != nil {
		return nil, err
	}

	// Get 'track to tab' links
	trackTabLinks, err := m.TracksTabsOps.GetTrackToTabLinks(trackID...)
	if err != nil {
		return nil, err
	}

	// Pluck the tabID's
	var tabIDs []uuid.UUID
	for _, trackTabLink := range trackTabLinks {
		tabIDs = append(tabIDs, trackTabLink.TabID)
	}

	// Get the tabs
	tabsMap, err := m.TabsOps.GetTabsAsMap(tabIDs...)

	// Map tabs to tracks
	for _, trackTabLink := range trackTabLinks {
		track := tracksMap[trackTabLink.TrackID]
		tab := tabsMap[trackTabLink.TabID]
		// Update track in tracksMap with new tab
		track.Tabs = append(track.Tabs, tab)
	}

	// Pluck the tracks from the map
	var tracks []*Track
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}

	// Done
	return tracks, nil
}

// GetTracksCascadingAsMap retrieves the tracks, with entity references, for the provided IDs, and creates a map where the key is the ID and the value the Track.
func (m *MappingService) GetTracksCascadingAsMap(trackID ...uuid.UUID) (map[uuid.UUID]*Track, error) {
	tracks, err := m.GetTracksCascading(trackID...)
	if err != nil {
		return nil, err
	}

	trackMap := make(map[uuid.UUID]*Track)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}

	return trackMap, nil
}
