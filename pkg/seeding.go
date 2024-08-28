package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
)

// Seeder represents all operations related to seeding.
type Seeder interface {
	Seed()
}

// SeedingAPI is responsible seeding data.
type SeedingAPI struct {
	*SvcManager
	*models.SeedConfig
}

// NewSeedingAPI instantiates a SeedingAPI.
func NewSeedingAPI(database logic.DbOperations, config *models.SeedConfig) Seeder {
	return &SeedingAPI{NewSvcManager(database), config}
}

// Seed seeds the database with the entries found in the provided models.SeedConfig.
func (s *SeedingAPI) Seed() {
	s.SeedInstruments()
	s.SeedDifficulties()
	s.SeedSources()
	s.SeedEndpoints()
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
