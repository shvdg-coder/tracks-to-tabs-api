package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TrackMapper represents operations related to track data mapping.
type TrackMapper interface {
	TracksToMap(tracks []*models.Track) map[uuid.UUID]*models.Track
	MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab) []*models.Track
	MapReferencesToTracks(trackMap map[uuid.UUID]*models.Track, references []*models.Reference) []*models.Track
}

// TrackSvc is responsible for mapping entities to tracks.
type TrackSvc struct {
	TrackMapper
}

// NewTrackSvc creates a new instance of ReferenceSvc.
func NewTrackSvc() TrackMapper {
	return &TrackSvc{}
}

// TracksToMap transforms a slice of tracks into a map where the key is the ID and the value the TrackEntry.
func (m *TrackSvc) TracksToMap(tracks []*models.Track) map[uuid.UUID]*models.Track {
	trackMap := make(map[uuid.UUID]*models.Track)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}
	return trackMap
}

// MapTabsToTracks adds the tabs to the tracks.
func (m *TrackSvc) MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab) []*models.Track {
	for _, link := range trackTabs {
		track := tracksMap[link.TrackID]
		tab := tabsMap[link.TabID]
		track.Tabs = append(track.Tabs, tab)
	}
	var tracks []*models.Track
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}
	return tracks
}

// MapReferencesToTracks maps references.Reference's to TrackEntry's.
func (m *TrackSvc) MapReferencesToTracks(trackMap map[uuid.UUID]*models.Track, references []*models.Reference) []*models.Track {
	for _, reference := range references {
		track := trackMap[reference.InternalID]
		track.References = append(track.References, reference)
	}
	var tracks []*models.Track
	for _, artist := range trackMap {
		tracks = append(tracks, artist)
	}
	return tracks
}
