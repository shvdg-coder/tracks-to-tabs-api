package pkg

import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"log"
	"strings"
)

// SeedOps represents all operations related to seeding.
type SeedOps interface {
	Seed()
}

// SeedingAPI is responsible for seeding data.
type SeedingAPI struct {
	SvcOps
	*SeedingConfig
	DummyOps
}

// NewSeedingAPI instantiates a SeedingAPI.
func NewSeedingAPI(svcManager SvcOps, config *SeedingConfig, dummies DummyOps) SeedOps {
	return &SeedingAPI{svcManager, config, dummies}
}

// Seed seeds the database with the entries found in the provided SeedingConfig.
func (s *SeedingAPI) Seed() {
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
	artists := s.SeedArtists()
	tracks := s.SeedTracks(artists)
	s.SeedTabs(tracks)
}

// SeedInstruments seeds the instruments table with the default instruments.
func (s *SeedingAPI) SeedInstruments() {
	err := s.InsertInstrumentEntries(s.SeedingConfig.Instruments...)
	if err != nil {
		log.Fatalf("Failed to insert instruments: %s", err.Error())
	}
}

// SeedDifficulties seeds the difficulties table with the default difficulties.
func (s *SeedingAPI) SeedDifficulties() {
	err := s.InsertDifficultyEntries(s.SeedingConfig.Difficulties...)
	if err != nil {
		log.Fatalf("Failed to insert difficulties: %s", err.Error())
	}
}

// SeedSources seeds the sources from the config file.
func (s *SeedingAPI) SeedSources() {
	err := s.InsertSourceEntries(s.SeedingConfig.Sources...)
	if err != nil {
		log.Fatalf("Failed to insert sources: %s", err.Error())
	}
}

// SeedEndpoints seeds the endpoints from the config file.
func (s *SeedingAPI) SeedEndpoints() {
	err := s.InsertEndpointEntries(s.SeedingConfig.Endpoints...)
	if err != nil {
		log.Fatalf("Failed to insert endpoints: %s", err.Error())
	}
}

// SeedArtists seeds the artists according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedArtists() []*models.ArtistEntry {
	dummyArtists := s.CreateArtists(s.Dummies.Artists)
	err := s.InsertArtistEntries(dummyArtists...)
	if err != nil {
		log.Fatalf("Failed to insert artists: %s", err.Error())
	}
	return dummyArtists
}

// SeedTracks seeds the tracks according to the dummy settings in the config file and returns their IDs.
func (s *SeedingAPI) SeedTracks(artists []*models.ArtistEntry) []*models.TrackEntry {
	dummyTracks := s.CreateTracks(s.Dummies.Artists.Tracks)
	err := s.InsertTrackEntries(dummyTracks...)
	if err != nil {
		log.Fatalf("Failed to insert tracks: %s", err.Error())
	}
	return dummyTracks
}

// SeedTabs seeds the tabs according to the dummy settings in the config file.
func (s *SeedingAPI) SeedTabs(tracks []*models.TrackEntry) []*models.TabEntry {
	dummyTabs := s.CreateTabs(s.Dummies.Artists.Tracks.Tabs)
	err := s.InsertTabEntries(dummyTabs...)
	if err != nil {
		log.Fatalf("Failed to insert tab entries: %v", err)
	}
	return dummyTabs
}

// formatName formats the provided name.
func (s *SeedingAPI) formatName(name string) string {
	return strings.ToLower(strings.Replace(name, " ", "-", -1))
}
