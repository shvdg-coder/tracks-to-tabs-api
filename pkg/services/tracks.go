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
	trackTabEntries, err := t.GetTrackToTabLinks(t.ExtractIDsFromTracks(tracks)...)
	if err != nil {
		return err
	}

	tabIDs := t.ExtractTabIDs(trackTabEntries)
	tabs, err := t.GetTabsCascading(tabIDs...)
	if err != nil {
		return err
	}

	tracksMap := t.TracksToMap(tracks)
	tabsMap := t.TabsToMap(tabs)
	t.MapTabsToTracks(tracksMap, tabsMap, trackTabEntries)

	return nil
}

// LoadReferences loads the references for the given tracks.
func (t *TrackSvc) LoadReferences(tracks ...*models.Track) error {
	references, err := t.GetReferences(t.ExtractIDsFromTracks(tracks)...)
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
