package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// MappingOperations represents operations related to track data mapping.
type MappingOperations interface {
	TracksToMap(tracks []*models.TrackEntry) map[uuid.UUID]*models.TrackEntry
	MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.TrackEntry, tabsMap map[uuid.UUID]*models.Tab) []*models.TrackEntry
	MapReferencesToTracks(trackMap map[uuid.UUID]*models.TrackEntry, references []*models.Reference) []*models.TrackEntry
}

// MappingService is responsible for mapping entities to tracks.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// TracksToMap transforms a slice of tracks into a map where the key is the ID and the value the TrackEntry.
func (m *MappingService) TracksToMap(tracks []*models.TrackEntry) map[uuid.UUID]*models.TrackEntry {
	trackMap := make(map[uuid.UUID]*models.TrackEntry)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}
	return trackMap
}

// MapTabsToTracks adds the tabs to the tracks.
func (m *MappingService) MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.TrackEntry, tabsMap map[uuid.UUID]*models.Tab) []*models.TrackEntry {
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
	var tracks []*models.TrackEntry
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}
	return tracks
}

// MapReferencesToTracks maps references.Reference's to TrackEntry's.
func (m *MappingService) MapReferencesToTracks(trackMap map[uuid.UUID]*models.TrackEntry, references []*models.Reference) []*models.TrackEntry {
	for _, reference := range references {
		track, ok := trackMap[reference.InternalID]
		if !ok {
			continue
		}
		track.References = append(track.References, reference)
	}
	var tracks []*models.TrackEntry
	for _, artist := range trackMap {
		tracks = append(tracks, artist)
	}
	return tracks
}
