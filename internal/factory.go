package internal

import (
	"errors"
	"fmt"
	faker "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// DummyFactory helps with creating dummy entities.
type DummyFactory struct {
	Sources      []*src.Source
	Instruments  []*inst.Instrument
	Difficulties []*diff.Difficulty
}

// NewFactory creates a new DummyFactory instance.
func NewFactory(sources []*src.Source, instruments []*inst.Instrument, difficulties []*diff.Difficulty) *DummyFactory {
	return &DummyFactory{
		Sources:      sources,
		Instruments:  instruments,
		Difficulties: difficulties,
	}
}

// GetRandomSource returns a random source that has the provided categories, from the DummyFactory's list of sources.
func (d *DummyFactory) GetRandomSource(categories ...string) (*src.Source, error) {
	var matchingSources []*src.Source
	for _, source := range d.Sources {
		if source.HasCategories(categories...) {
			matchingSources = append(matchingSources, source)
		}
	}
	if len(matchingSources) == 0 {
		return nil, errors.New(fmt.Sprintf("A source with the categories '%s' does not exist", categories))
	}
	return matchingSources[faker.Number(0, len(matchingSources)-1)], nil
}

// GetRandomInstrument returns a random instrument from the DummyFactory's list of instruments.
func (d *DummyFactory) GetRandomInstrument() *inst.Instrument {
	return d.Instruments[faker.Number(0, len(d.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the DummyFactory's list of difficulties.
func (d *DummyFactory) GetRandomDifficulty() *diff.Difficulty {
	return d.Difficulties[faker.Number(0, len(d.Difficulties)-1)]
}

// CreateReferenceID creates a new reference ID based on the provided internal ID and category.
func (d *DummyFactory) CreateReferenceID(internalId uuid.UUID, category string) *references.Reference {
	sourceId, _ := d.GetRandomSource(category)
	return references.NewReference(
		internalId,
		sourceId.ID,
		category,
		"ID",
		faker.UUID(),
	)
}

// CreateArtists creates a specified amount of dummy artists.
func (d *DummyFactory) CreateArtists(artists *Artists) []*art.Artist {
	dummyArtists := make([]*art.Artist, artists.randomAmount())
	for i := range dummyArtists {
		dummyArtists[i] = d.CreateArtist(artists.Tracks)
	}
	return dummyArtists
}

// CreateArtist creates a dummy artist with a random name and tracks.
func (d *DummyFactory) CreateArtist(tracks *Tracks) *art.Artist {
	return art.NewArtist(
		faker.HipsterWord(),
		art.WithTracks(d.CreateTracks(tracks)))
}

// CreateTracks creates a specified amount of dummy tracks.
func (d *DummyFactory) CreateTracks(tracks *Tracks) []*trk.Track {
	dummyTracks := make([]*trk.Track, tracks.randomAmount())
	for i := range dummyTracks {
		dummyTracks[i] = d.CreateTrack(tracks.Tabs)
	}
	return dummyTracks
}

// CreateTrack creates a dummy track with a random title, duration, and tabs.
func (d *DummyFactory) CreateTrack(tabs *Tabs) *trk.Track {
	return trk.NewTrack(
		faker.HipsterSentence(faker.Number(1, 6)),
		uint(faker.Number(10000, 3000000)), // 1 to 5 minutes
		trk.WithTabs(d.CreateTabs(tabs.randomAmount())))
}

// CreateTabs creates a specified amount of dummy tabs.
func (d *DummyFactory) CreateTabs(amount uint) []*tabs.Tab {
	dummyTabs := make([]*tabs.Tab, amount)
	for i := range dummyTabs {
		dummyTabs[i] = d.CreateTab()
	}
	return dummyTabs
}

// CreateTab creates a new dummy tab with a random instrument, difficulty, and description
func (d *DummyFactory) CreateTab() *tabs.Tab {
	return tabs.NewTab(d.GetRandomInstrument(), d.GetRandomDifficulty(), faker.Name())
}
