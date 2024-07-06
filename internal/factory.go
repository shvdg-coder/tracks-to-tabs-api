package internal

import (
	"errors"
	"fmt"
	faker "github.com/brianvoe/gofakeit/v7"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// Factory helps with creating entities.
type Factory struct {
	Sources      []*src.Source
	Instruments  []*inst.Instrument
	Difficulties []*diff.Difficulty
}

// NewFactory creates a new Seeding.Dummies instance.
func NewFactory(instruments []*inst.Instrument, difficulties []*diff.Difficulty) *Factory {
	return &Factory{
		Instruments:  instruments,
		Difficulties: difficulties,
	}
}

// GetRandomSource returns a random source that has the provided categories, from the Factory's list of sources.
func (f *Factory) GetRandomSource(categories ...string) (*src.Source, error) {
	var matchingSources []*src.Source
	for _, source := range f.Sources {
		if source.HasCategories(categories...) {
			matchingSources = append(matchingSources, source)
		}
	}
	if len(matchingSources) == 0 {
		return nil, errors.New(fmt.Sprintf("A source with the categories '%s' does not exist", categories))
	}
	return matchingSources[faker.Number(0, len(matchingSources)-1)], nil
}

// GetRandomInstrument returns a random instrument from the Factory's list of instruments.
func (f *Factory) GetRandomInstrument() *inst.Instrument {
	return f.Instruments[faker.Number(0, len(f.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the Factory's list of difficulties.
func (f *Factory) GetRandomDifficulty() *diff.Difficulty {
	return f.Difficulties[faker.Number(0, len(f.Difficulties)-1)]
}

// CreateDummyArtists creates a specified amount of dummy artists.
func (f *Factory) CreateDummyArtists(artists *Artists) []*art.Artist {
	dummyArtists := make([]*art.Artist, artists.randomAmount())
	for i := range dummyArtists {
		dummyArtists[i] = f.CreateDummyArtist(artists.Tracks)
	}
	return dummyArtists
}

// CreateDummyArtist creates a dummy artist with a random name and tracks.
func (f *Factory) CreateDummyArtist(tracks *Tracks) *art.Artist {
	return art.NewArtist(
		faker.HipsterWord(),
		art.WithTracks(f.CreateDummyTracks(tracks)))
}

// CreateDummyTracks creates a specified amount of dummy tracks.
func (f *Factory) CreateDummyTracks(tracks *Tracks) []*trk.Track {
	dummyTracks := make([]*trk.Track, tracks.randomAmount())
	for i := range dummyTracks {
		dummyTracks[i] = f.createDummyTrack(tracks.Tabs)
	}
	return dummyTracks
}

// createDummyTrack creates a dummy track with a random title, duration, and tabs.
func (f *Factory) createDummyTrack(tabs *Tabs) *trk.Track {
	return trk.NewTrack(
		faker.HipsterSentence(faker.Number(1, 6)),
		uint(faker.Number(10000, 3000000)), // 1 to 5 minutes
		trk.WithTabs(f.createDummyTabs(tabs.randomAmount())))
}

// createDummyTabs creates a specified amount of dummy tabs.
func (f *Factory) createDummyTabs(amount uint) []*tabs.Tab {
	dummyTabs := make([]*tabs.Tab, amount)
	for i := range dummyTabs {
		dummyTabs[i] = f.CreateDummyTab()
	}
	return dummyTabs
}

// CreateDummyTab creates a new dummy tab with a random instrument, difficulty, and description
func (f *Factory) CreateDummyTab() *tabs.Tab {
	return tabs.NewTab(f.GetRandomInstrument(), f.GetRandomDifficulty(), faker.Name())
}
