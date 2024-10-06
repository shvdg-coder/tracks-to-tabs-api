package services

import (
	"encoding/base64"
	faker "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
)

// DummyOps represents all operations related to dummy data.
type DummyOps interface {
	GetRandomSource(category string) *models.SourceEntry
	CreateRandomUUID() string

	CreateReference(internalID uuid.UUID, sourceID uint, referenceType, referenceCategory, value string) *models.ReferenceEntry
	CreateArtists(config *models.ArtistConfig) []*models.ArtistEntry
	CreateArtist() *models.ArtistEntry
	CreateTracks(config *models.TrackConfig) []*models.TrackEntry
	CreateTrack() *models.TrackEntry
	CreateTabs(config *models.TabConfig) []*models.TabEntry
	CreateTab() *models.TabEntry
}

// DummySvc is responsible for dummy data.
type DummySvc struct {
	*SvcManager
	Sources      []*models.SourceEntry
	Instruments  []*models.InstrumentEntry
	Difficulties []*models.DifficultyEntry
}

// NewDummySvc instantiates a DummySvc.
func NewDummySvc(svcManager *SvcManager, sources []*models.SourceEntry, instruments []*models.InstrumentEntry, difficulties []*models.DifficultyEntry) DummyOps {
	return &DummySvc{
		svcManager,
		sources,
		instruments,
		difficulties}
}

// GetRandomSource returns a random source that has the provided category, from the DummyService's list of sources.
func (d *DummySvc) GetRandomSource(category string) *models.SourceEntry {
	var matchingSources []*models.SourceEntry
	for _, source := range d.Sources {
		if source.HasCategory(category) {
			matchingSources = append(matchingSources, source)
		}
	}
	if len(matchingSources) == 0 {
		return nil
	}
	return matchingSources[faker.Number(0, len(matchingSources)-1)]
}

// GetRandomInstrument returns a random instrument from the DummySvc list of instruments.
func (d *DummySvc) GetRandomInstrument() *models.InstrumentEntry {
	return d.Instruments[faker.Number(0, len(d.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the DummySvc list of difficulties.
func (d *DummySvc) GetRandomDifficulty() *models.DifficultyEntry {
	return d.Difficulties[faker.Number(0, len(d.Difficulties)-1)]
}

// CreateRandomUUID creates a random uuid as a string.
func (d *DummySvc) CreateRandomUUID() string {
	return faker.UUID()
}

// CreateReference creates a new models.Reference using the provided values.
func (d *DummySvc) CreateReference(internalID uuid.UUID, sourceID uint, referenceType, referenceCategory, value string) *models.ReferenceEntry {
	return &models.ReferenceEntry{
		InternalID: internalID,
		SourceID:   sourceID,
		Category:   referenceCategory,
		Type:       referenceType,
		Reference:  value,
	}
}

// CreateArtists creates a specified amount of dummy artists.
func (d *DummySvc) CreateArtists(config *models.ArtistConfig) []*models.ArtistEntry {
	dummyArtists := make([]*models.ArtistEntry, config.RandomAmount())
	for i := range dummyArtists {
		dummyArtists[i] = d.CreateArtist()
	}
	return dummyArtists
}

// CreateArtist creates a dummy artist with a random name and tracks.
func (d *DummySvc) CreateArtist() *models.ArtistEntry {
	coverImg := faker.ImageJpeg(750, 750)
	bannerImg := faker.ImageJpeg(2660, 1140)
	return &models.ArtistEntry{
		ID:     uuid.New(),
		Name:   faker.HipsterWord(),
		Cover:  base64.StdEncoding.EncodeToString(coverImg),
		Banner: base64.StdEncoding.EncodeToString(bannerImg),
	}
}

// CreateTracks creates a specified amount of dummy tracks.
func (d *DummySvc) CreateTracks(config *models.TrackConfig) []*models.TrackEntry {
	dummyTracks := make([]*models.TrackEntry, config.RandomAmount())
	for i := range dummyTracks {
		dummyTracks[i] = d.CreateTrack()
	}
	return dummyTracks
}

// CreateTrack creates a dummy track with a random title, duration, and tabs.
func (d *DummySvc) CreateTrack() *models.TrackEntry {
	coverImg := faker.ImageJpeg(750, 750)
	return &models.TrackEntry{
		ID:       uuid.New(),
		Title:    faker.HipsterSentence(faker.Number(1, 6)),
		Cover:    base64.StdEncoding.EncodeToString(coverImg),
		Duration: uint(faker.Number(10000, 3000000)),
	}
}

// CreateTabs creates a specified amount of dummy tabs.
func (d *DummySvc) CreateTabs(config *models.TabConfig) []*models.TabEntry {
	dummyTabs := make([]*models.TabEntry, config.RandomAmount())
	for i := range dummyTabs {
		dummyTabs[i] = d.CreateTab()
	}
	return dummyTabs
}

// CreateTab creates a new dummy tab with a random instrument, difficulty, and description
func (d *DummySvc) CreateTab() *models.TabEntry {
	return &models.TabEntry{
		ID:           uuid.New(),
		InstrumentID: d.GetRandomInstrument().ID,
		DifficultyID: d.GetRandomDifficulty().ID,
		Description:  faker.Name(),
	}
}
