package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// TrackOps represent all operations related to tracks.
type TrackOps interface {
	schemas.TrackSchema
	data.TrackData
	mappers.TrackMapper
	GetTracks(trackID ...uuid.UUID) ([]*models.Track, error)
	GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error)
	ExtractIDsFromTracks(tracks []*models.Track) []uuid.UUID
	CollectTabs(tracks []*models.Track) []*models.Tab
}

// TrackSvc is responsible for managing and retrieving tracks.
type TrackSvc struct {
	schemas.TrackSchema
	data.TrackData
	mappers.TrackMapper
	TrackTabOps
	TabOps
	ReferenceOps
}

// NewTrackSvc instantiates a TrackSvc.
func NewTrackSvc(schema schemas.TrackSchema, data data.TrackData, mapper mappers.TrackMapper, trackTabs TrackTabOps, tabs TabOps, references ReferenceOps) TrackOps {
	return &TrackSvc{
		TrackSchema:  schema,
		TrackData:    data,
		TrackMapper:  mapper,
		TrackTabOps:  trackTabs,
		TabOps:       tabs,
		ReferenceOps: references,
	}
}

// GetTracks retrieves tracks, without entity references, for the provided IDs.
func (t *TrackSvc) GetTracks(trackID ...uuid.UUID) ([]*models.Track, error) {
	trackEntries, err := t.GetTrackEntries(trackID...)
	if err != nil {
		return nil, err
	}

	tracks := t.TrackEntriesToTracks(trackEntries)

	return tracks, nil
}

// GetTracksCascading retrieves tracks, with entity references, for the provided IDs.
func (t *TrackSvc) GetTracksCascading(trackID ...uuid.UUID) ([]*models.Track, error) {
	tracks, err := t.GetTracks(trackID...)
	if err != nil {
		return nil, err
	}

	err = t.LoadTabs(tracks...)
	if err != nil {
		return nil, err
	}

	err = t.LoadReferences(tracks...)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

// LoadTabs loads the models.Tab's for the given models.Track's.
func (t *TrackSvc) LoadTabs(tracks ...*models.Track) error {
	trackIDs := t.ExtractIDsFromTracks(tracks)
	trackTabEntries, err := t.GetTrackToTabLinks(trackIDs...)
	if err != nil {
		return err
	}

	_, tabIDs := t.ExtractIDsFromTrackTabEntries(trackTabEntries)
	tabs, err := t.GetTabsCascading(tabIDs...)
	if err != nil {
		return err
	}

	tracksMap := t.TracksToMap(tracks)
	tabsMap := t.TabsToMap(tabs)

	t.MapTracksToTabs(tabsMap, tracksMap, trackTabEntries)
	t.MapTabsToTracks(tracksMap, tabsMap, trackTabEntries)

	return nil
}

// LoadReferences loads the references for the given tracks.
func (t *TrackSvc) LoadReferences(tracks ...*models.Track) error {
	trackIDs := t.ExtractIDsFromTracks(tracks)
	references, err := t.GetReferencesCascading(trackIDs...)
	if err != nil {
		return err
	}

	tracksMap := t.TracksToMap(tracks)
	t.MapReferencesToTracks(tracksMap, references)

	return nil
}

// ExtractIDsFromTracks retrieves the ID's from the models.Track's.
func (t *TrackSvc) ExtractIDsFromTracks(tracks []*models.Track) []uuid.UUID {
	trackIDs := make([]uuid.UUID, len(tracks))
	for i, track := range tracks {
		trackIDs[i] = track.ID
	}
	return trackIDs
}

// CollectTabs pluck the models.Tab's from each of the models.Track's, and returns them.
func (t *TrackSvc) CollectTabs(tracks []*models.Track) []*models.Tab {
	tabs := make([]*models.Tab, 0)
	for _, track := range tracks {
		tabs = append(tabs, track.Tabs...)
	}
	return tabs
}
