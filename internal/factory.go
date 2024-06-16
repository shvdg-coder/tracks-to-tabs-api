package internal

import (
	gofakeit "github.com/brianvoe/gofakeit/v7"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// Factory helps with creating entities.
type Factory struct {
	Instruments  []*inst.Instrument
	Difficulties []*diff.Difficulty
}

// NewFactory creates a new Factory instance.
func NewFactory(instruments []*inst.Instrument, difficulties []*diff.Difficulty) *Factory {
	return &Factory{
		Instruments:  instruments,
		Difficulties: difficulties,
	}
}

// GetRandomInstrument returns a random instrument from the Factory's list of instruments.
func (f *Factory) GetRandomInstrument() *inst.Instrument {
	return f.Instruments[gofakeit.Number(0, len(f.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the Factory's list of difficulties.
func (f *Factory) GetRandomDifficulty() *diff.Difficulty {
	return f.Difficulties[gofakeit.Number(0, len(f.Difficulties)-1)]
}

// CreateDummyArtists creates a specified amount of dummy artists.
func (f *Factory) CreateDummyArtists(amount uint) []*art.Artist {
	dummyArtists := make([]*art.Artist, amount)
	for i := range dummyArtists {
		dummyArtists[i] = f.CreateDummyArtist()
	}
	return dummyArtists
}

// CreateDummyArtist creates a dummy artist with a random name and tracks.
func (f *Factory) CreateDummyArtist() *art.Artist {
	return art.NewArtist(
		gofakeit.HipsterWord(),
		art.WithTracks(f.CreateDummyTracks(uint(gofakeit.Number(0, 20)))))
}

// CreateDummyTracks creates a specified amount of dummy tracks.
func (f *Factory) CreateDummyTracks(amount uint) []*trk.Track {
	dummyTracks := make([]*trk.Track, amount)
	for i := range dummyTracks {
		dummyTracks[i] = f.createDummyTrack()
	}
	return dummyTracks
}

// createDummyTrack creates a dummy track with a random title, duration, and tabs.
func (f *Factory) createDummyTrack() *trk.Track {
	return trk.NewTrack(
		gofakeit.HipsterSentence(gofakeit.Number(1, 6)),
		uint(gofakeit.Number(10000, 3000000)), // 1 to 5 minutes
		trk.WithTabs(f.createDummyTabs(uint(gofakeit.Number(0, 4)))))
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
	return tabs.NewTab(f.GetRandomInstrument(), f.GetRandomDifficulty(), gofakeit.Name())
}
