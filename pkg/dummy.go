package pkg

import (
	faker "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
)

// DummyOps represents all operations related to dummy data.
type DummyOps interface {
	GetRandomSource(category string) *models.SourceEntry
	CreateRandomUUID() string

	CreateReference(internalID uuid.UUID, sourceID uint, referenceType, referenceCategory, value string) *models.ReferenceEntry
	CreateArtists(config *models.ArtistsConfig) []*models.ArtistEntry
	CreateArtist() *models.ArtistEntry
	CreateTracks(config *models.TracksConfig) []*models.TrackEntry
	CreateTrack() *models.TrackEntry
	CreateTabs(config *models.TabsConfig) []*models.TabEntry
	CreateTab() *models.TabEntry
}

// DummyAPI is responsible dummy data.
type DummyAPI struct {
	*SvcManager
	Sources      []*models.SourceEntry
	Instruments  []*models.InstrumentEntry
	Difficulties []*models.DifficultyEntry
}

// NewDummyAPI instantiates a DummyAPI.
func NewDummyAPI(database logic.DbOperations, sources []*models.SourceEntry, instruments []*models.InstrumentEntry, difficulties []*models.DifficultyEntry) DummyOps {
	return &DummyAPI{
		NewSvcManager(database),
		sources,
		instruments,
		difficulties}
}

// GetRandomSource returns a random source that has the provided category, from the DummyService's list of sources.
func (d *DummyAPI) GetRandomSource(category string) *models.SourceEntry {
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

// GetRandomInstrument returns a random instrument from the DummyAPI list of instruments.
func (d *DummyAPI) GetRandomInstrument() *models.InstrumentEntry {
	return d.Instruments[faker.Number(0, len(d.Instruments)-1)]
}

// GetRandomDifficulty returns a random difficulty from the DummyAPI list of difficulties.
func (d *DummyAPI) GetRandomDifficulty() *models.DifficultyEntry {
	return d.Difficulties[faker.Number(0, len(d.Difficulties)-1)]
}

// CreateRandomUUID creates a random uuid as a string.
func (d *DummyAPI) CreateRandomUUID() string {
	return faker.UUID()
}

// CreateReference creates a new models.Reference using the provided values.
func (d *DummyAPI) CreateReference(internalID uuid.UUID, sourceID uint, referenceType, referenceCategory, value string) *models.ReferenceEntry {
	return &models.ReferenceEntry{
		InternalID: internalID,
		SourceID:   sourceID,
		Category:   referenceCategory,
		Type:       referenceType,
		Reference:  value,
	}
}

// CreateArtists creates a specified amount of dummy artists.
func (d *DummyAPI) CreateArtists(config *models.ArtistsConfig) []*models.ArtistEntry {
	dummyArtists := make([]*models.ArtistEntry, config.RandomAmount())
	for i := range dummyArtists {
		dummyArtists[i] = d.CreateArtist()
	}
	return dummyArtists
}

// CreateArtist creates a dummy artist with a random name and tracks.
func (d *DummyAPI) CreateArtist() *models.ArtistEntry {
	return &models.ArtistEntry{
		ID:   uuid.New(),
		Name: faker.HipsterWord(),
	}
}

// CreateTracks creates a specified amount of dummy tracks.
func (d *DummyAPI) CreateTracks(config *models.TracksConfig) []*models.TrackEntry {
	dummyTracks := make([]*models.TrackEntry, config.RandomAmount())
	for i := range dummyTracks {
		dummyTracks[i] = d.CreateTrack()
	}
	return dummyTracks
}

// CreateTrack creates a dummy track with a random title, duration, and tabs.
func (d *DummyAPI) CreateTrack() *models.TrackEntry {
	return &models.TrackEntry{
		ID:       uuid.New(),
		Title:    faker.HipsterSentence(faker.Number(1, 6)),
		Duration: uint(faker.Number(10000, 3000000)),
	}
}

// CreateTabs creates a specified amount of dummy tabs.
func (d *DummyAPI) CreateTabs(config *models.TabsConfig) []*models.TabEntry {
	dummyTabs := make([]*models.TabEntry, config.RandomAmount())
	for i := range dummyTabs {
		dummyTabs[i] = d.CreateTab()
	}
	return dummyTabs
}

// CreateTab creates a new dummy tab with a random instrument, difficulty, and description
func (d *DummyAPI) CreateTab() *models.TabEntry {
	return &models.TabEntry{
		ID:           uuid.New(),
		InstrumentID: d.GetRandomInstrument().ID,
		DifficultyID: d.GetRandomDifficulty().ID,
		Description:  faker.Name(),
	}
}
