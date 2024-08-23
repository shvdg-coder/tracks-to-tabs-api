package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TrackMapper represents operations related to track data mapping.
type TrackMapper interface {
	TrackEntriesToTracks(trackEntries []*models.TrackEntry) []*models.Track
	TracksToMap(tracks []*models.Track) map[uuid.UUID]*models.Track
	MapToTracks(tracksMap map[uuid.UUID]*models.Track) []*models.Track
	MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab) map[uuid.UUID]*models.Track
	MapReferencesToTracks(tracksMap map[uuid.UUID]*models.Track, references []*models.Reference) map[uuid.UUID]*models.Track
}

// TrackSvc is responsible for mapping entities to tracks.
type TrackSvc struct {
	TrackMapper
}

// NewTrackSvc creates a new instance of ReferenceSvc.
func NewTrackSvc() TrackMapper {
	return &TrackSvc{}
}

// TrackEntriesToTracks transforms the models.TrackEntry's to models.Track's, with default values where needed.
func (t *TrackSvc) TrackEntriesToTracks(trackEntries []*models.TrackEntry) []*models.Track {
	tracks := make([]*models.Track, len(trackEntries))
	for i, trackEntry := range trackEntries {
		tracks[i] = &models.Track{
			TrackEntry: trackEntry,
			Artist:     &models.Artist{},
			Tabs:       make([]*models.Tab, 0),
			References: make([]*models.Reference, 0),
		}
	}
	return tracks
}

// TracksToMap transforms a slice of models.Track's into a map where the key is the ID and the value the models.TrackEntry.
func (t *TrackSvc) TracksToMap(tracks []*models.Track) map[uuid.UUID]*models.Track {
	trackMap := make(map[uuid.UUID]*models.Track)
	for _, track := range tracks {
		trackMap[track.ID] = track
	}
	return trackMap
}

// MapToTracks transforms a map of Tracks into a slice of models.Track's.
func (a *ArtistSvc) MapToTracks(tracksMap map[uuid.UUID]*models.Track) []*models.Track {
	tracks := make([]*models.Track, len(tracksMap))
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}
	return tracks
}

// MapTabsToTracks adds the tabs to the tracks.
func (t *TrackSvc) MapTabsToTracks(trackTabs []*models.TrackTabEntry, tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab) map[uuid.UUID]*models.Track {
	for _, link := range trackTabs {
		track := tracksMap[link.TrackID]
		tab := tabsMap[link.TabID]
		track.Tabs = append(track.Tabs, tab)
	}
	return tracksMap
}

// MapReferencesToTracks maps references.Reference's to TrackEntry's.
func (t *TrackSvc) MapReferencesToTracks(tracksMap map[uuid.UUID]*models.Track, references []*models.Reference) map[uuid.UUID]*models.Track {
	for _, reference := range references {
		track := tracksMap[reference.InternalID]
		track.References = append(track.References, reference)
	}
	return tracksMap
}
