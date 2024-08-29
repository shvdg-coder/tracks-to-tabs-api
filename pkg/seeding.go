package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"strings"
)

// Seeder represents all operations related to seeding.
type Seeder interface {
	Seed()
}

// SeedingAPI is responsible for seeding data.
type SeedingAPI struct {
	*SvcManager
	*models.SeedConfig
	DummyOps
}

// NewSeedingAPI instantiates a SeedingAPI.
func NewSeedingAPI(database logic.DbOperations, config *models.SeedConfig, dummies DummyOps) Seeder {
	return &SeedingAPI{NewSvcManager(database), config, dummies}
}

// Seed seeds the database with the entries found in the provided models.SeedConfig.
func (s *SeedingAPI) Seed() {
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
	s.SeedArtists()
	s.SeedTracks()
	s.SeedTabs()
}

// SeedInstruments seeds the instruments table with the default instruments.
func (s *SeedingAPI) SeedInstruments() {
	s.InsertInstrumentEntries(s.SeedConfig.Instruments...)
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedingAPI) SeedDifficulties() {
	s.InsertDifficultyEntries(s.SeedConfig.Difficulties...)
}

// SeedSources seeds the sources from the config file.
func (s *SeedingAPI) SeedSources() {
	s.InsertSourceEntries(s.SeedConfig.Sources...)
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedingAPI) SeedEndpoints() {
	s.InsertEndpointEntries(s.SeedConfig.Endpoints...)
}

// SeedArtists seeds the artists according to the dummy settings in the config file.
func (s *SeedingAPI) SeedArtists() {
	dummyArtists := s.CreateArtists(s.Dummies.Artists)
	for _, artist := range dummyArtists {
		s.InsertArtistEntry(artist)

		sourceMusic := s.GetRandomSource(CategoryMusic)
		artistIDRef := s.CreateReference(artist.ID, sourceMusic.ID, TypeID, CategoryArtist, s.CreateRandomUUID())
		s.InsertReferenceEntry(artistIDRef)

		sourceTabs := s.GetRandomSource(CategoryTabs)
		artistNameRef := s.CreateReference(artist.ID, sourceTabs.ID, TypeName, CategoryArtist, s.formatName(artist.Name))
		s.InsertReferenceEntry(artistNameRef)
	}
}

// SeedTracks seeds the tracks according to the dummy settings in the config file.
func (s *SeedingAPI) SeedTracks() {
	dummyTracks := s.CreateTracks(s.Dummies.Artists.Tracks)
	for _, track := range dummyTracks {
		s.InsertTrackEntry(track)

		sourceMusic := s.GetRandomSource(CategoryMusic)
		trackIDRef := s.CreateReference(track.ID, sourceMusic.ID, TypeID, CategoryTrack, s.CreateRandomUUID())
		s.InsertReferenceEntry(trackIDRef)

		sourceTabs := s.GetRandomSource(CategoryTabs)
		trackNameRef := s.CreateReference(track.ID, sourceTabs.ID, TypeName, CategoryTrack, s.formatName(track.Title))
		s.InsertReferenceEntry(trackNameRef)
	}
}

// SeedTabs seeds the tabs according to the dummy settings in the config file.
func (s *SeedingAPI) SeedTabs() {
	dummyTabs := s.CreateTabs(s.Dummies.Artists.Tracks.Tabs)
	for _, tab := range dummyTabs {
		s.InsertTabEntry(tab)

		sourceTabs := s.GetRandomSource(CategoryTabs)

		tabIDRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeID, CategoryTab, s.CreateRandomUUID())
		s.InsertReferenceEntry(tabIDRef)

		tabNameRef := s.CreateReference(tab.ID, sourceTabs.ID, TypeName, CategoryTab, s.formatName(tab.Description))
		s.InsertReferenceEntry(tabNameRef)
	}
}

// formatName formats the provided name.
func (s *SeedingAPI) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
