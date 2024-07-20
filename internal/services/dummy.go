package services

import (
	"errors"
	"fmt"
	faker "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// DummyService helps with creating dummy entities.
type DummyService struct {
	Sources      []*src.Source
	Instruments  []*inst.Instrument
	Difficulties []*diff.Difficulty
}

// NewDummyService creates a new DummyService instance.
func NewDummyService(sources []*src.Source, instruments []*inst.Instrument, difficulties []*diff.Difficulty) *DummyService {
	return &DummyService{
		Sources:      sources,
		Instruments:  instruments,
		Difficulties: difficulties,
	}
}

// GetRandomSource returns a random source that has the provided category, from the DummyService's list of sources.
func (d *DummyService) GetRandomSource(category string) (*src.Source, error) {
	var matchingSources []*src.Source
	for _, source := range d.Sources {
		if source.HasCategory(category) {
			matchingSources = append(matchingSources, source)
		}
	}
	if len(matchingSources) == 0 {
		return nil, errors.New(fmt.Sprintf("A source with the category '%s' does not exist", category))
	}
	return matchingSources[faker.Number(0, len(matchingSources)-1)], nil
}

// GetRandomInstrument returns a random instrument from the DummyService's list of instruments.
func (d *DummyService) GetRandomInstrument() *inst.Instrument {
	return d.Instruments[faker.Number(0, len(d.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the DummyService's list of difficulties.
func (d *DummyService) GetRandomDifficulty() *diff.Difficulty {
	return d.Difficulties[faker.Number(0, len(d.Difficulties)-1)]
}

// CreateReferenceID creates a new reference ID, based on the provided internal ID and categories.
func (d *DummyService) CreateReferenceID(internalId uuid.UUID, sourceCategory, referenceCategory string) *references.Reference {
	sourceId, _ := d.GetRandomSource(sourceCategory)
	return references.NewReference(
		internalId,
		sourceId.ID,
		referenceCategory,
		"ID",
		faker.UUID(),
	)
}

// CreateArtists creates a specified amount of dummy artists.
func (d *DummyService) CreateArtists(artists *internal.ArtistsConfig) []*art.Artist {
	dummyArtists := make([]*art.Artist, artists.RandomAmount())
	for i := range dummyArtists {
		dummyArtists[i] = d.CreateArtist(artists.Tracks)
	}
	return dummyArtists
}

// CreateArtist creates a dummy artist with a random name and tracks.
func (d *DummyService) CreateArtist(tracks *internal.TracksConfig) *art.Artist {
	return art.NewArtist(
		faker.HipsterWord(),
		art.WithTracks(d.CreateTracks(tracks)))
}

// CreateTracks creates a specified amount of dummy tracks.
func (d *DummyService) CreateTracks(tracks *internal.TracksConfig) []*trk.Track {
	dummyTracks := make([]*trk.Track, tracks.RandomAmount())
	for i := range dummyTracks {
		dummyTracks[i] = d.CreateTrack(tracks.Tabs)
	}
	return dummyTracks
}

// CreateTrack creates a dummy track with a random title, duration, and tabs.
func (d *DummyService) CreateTrack(tabs *internal.TabsConfig) *trk.Track {
	return trk.NewTrack(
		faker.HipsterSentence(faker.Number(1, 6)),
		uint(faker.Number(10000, 3000000)), // 1 to 5 minutes
		trk.WithTabs(d.CreateTabs(tabs.RandomAmount())))
}

// CreateTabs creates a specified amount of dummy tabs.
func (d *DummyService) CreateTabs(amount uint) []*tabs.Tab {
	dummyTabs := make([]*tabs.Tab, amount)
	for i := range dummyTabs {
		dummyTabs[i] = d.CreateTab()
	}
	return dummyTabs
}

// CreateTab creates a new dummy tab with a random instrument, difficulty, and description
func (d *DummyService) CreateTab() *tabs.Tab {
	return tabs.NewTab(d.GetRandomInstrument(), d.GetRandomDifficulty(), faker.Name())
}
