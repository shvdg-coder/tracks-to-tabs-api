package tracks

import (
	"github.com/google/uuid"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// MappingOperations represents operations related to track data mapping.
type MappingOperations interface {
	ToMap(tracks []*Track) map[uuid.UUID]*Track
	MapTabsToTracks(trackTabs []*trktab.TrackTab, tracksMap map[uuid.UUID]*Track, tabsMap map[uuid.UUID]*tbs.Tab) []*Track
}

// MappingService is responsible for mapping entities to tracks.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// ToMap transforms a slice of tracks into a map where the key is the ID and the value the Track.
func (m *MappingService) ToMap(tracks []*Track) map[uuid.UUID]*Track {
	trackMap := make(map[uuid.UUID]*Track)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}
	return trackMap
}

// MapTabsToTracks adds the tabs to the tracks.
func (m *MappingService) MapTabsToTracks(trackTabs []*trktab.TrackTab, tracksMap map[uuid.UUID]*Track, tabsMap map[uuid.UUID]*tbs.Tab) []*Track {
	for _, link := range trackTabs {
		track, ok := tracksMap[link.TrackID]
		if !ok {
			continue
		}
		tab, ok := tabsMap[link.TabID]
		if !ok {
			continue
		}
		track.Tabs = append(track.Tabs, tab)
	}
	var tracks []*Track
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}
	return tracks
}
