package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
)

// TrackMapper represents operations related to track data mapping.
type TrackMapper interface {
	TrackEntriesToTracks(trackEntries []*models.TrackEntry) []*models.Track
	TracksToMap(tracks []*models.Track) map[uuid.UUID]*models.Track
	MapToTracks(tracksMap map[uuid.UUID]*models.Track) []*models.Track
	MapArtistsToTracks(tracksMap map[uuid.UUID]*models.Track, artistsMap map[uuid.UUID]*models.Artist, artistTracks []*models.ArtistTrackEntry) map[uuid.UUID]*models.Track
	MapTabsToTracks(tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab, trackTabs []*models.TrackTabEntry) map[uuid.UUID]*models.Track
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

// TrackEntriesToTracks transforms the models.TrackEntry's to models.Track's.
func (t *TrackSvc) TrackEntriesToTracks(trackEntries []*models.TrackEntry) []*models.Track {
	tracks := make([]*models.Track, len(trackEntries))
	for i, trackEntry := range trackEntries {
		tracks[i] = &models.Track{TrackEntry: trackEntry}
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

// MapToTracks transforms a map of models.Track into a slice of models.Track's.
func (t *TrackSvc) MapToTracks(tracksMap map[uuid.UUID]*models.Track) []*models.Track {
	tracks := make([]*models.Track, 0)
	for _, track := range tracksMap {
		tracks = append(tracks, track)
	}
	return tracks
}

// MapArtistsToTracks maps the models.Artist to the models.Track, by updating the provided models.Track's map and returning it.
func (t *TrackSvc) MapArtistsToTracks(tracksMap map[uuid.UUID]*models.Track, artistsMap map[uuid.UUID]*models.Artist, artistTracks []*models.ArtistTrackEntry) map[uuid.UUID]*models.Track {
	for _, artistTrack := range artistTracks {
		artist := artistsMap[artistTrack.ArtistID]
		track := tracksMap[artistTrack.TrackID]
		track.Artist = artist
	}
	return tracksMap
}

// MapTabsToTracks adds the models.Tab's to the models.Track, by updating the provided models.Track's map and returning it.
func (t *TrackSvc) MapTabsToTracks(tracksMap map[uuid.UUID]*models.Track, tabsMap map[uuid.UUID]*models.Tab, trackTabs []*models.TrackTabEntry) map[uuid.UUID]*models.Track {
	for _, trackTab := range trackTabs {
		track := tracksMap[trackTab.TrackID]
		tab := tabsMap[trackTab.TabID]
		track.Tabs = append(track.Tabs, tab)
	}
	return tracksMap
}

// MapReferencesToTracks adds the models.Reference's to the models.Track, by updating the provided models.Track's map and returning it.
func (t *TrackSvc) MapReferencesToTracks(tracksMap map[uuid.UUID]*models.Track, references []*models.Reference) map[uuid.UUID]*models.Track {
	for _, reference := range references {
		track := tracksMap[reference.InternalID]
		track.References = append(track.References, reference)
	}
	return tracksMap
}
