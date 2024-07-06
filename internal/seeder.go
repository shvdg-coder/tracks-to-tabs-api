package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Seeder helps with deleting data from the database
type Seeder struct {
	Seeding *Seeding
	API     *api.API
	Factory *Factory
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(seeding *Seeding, api *api.API) *Seeder {
	return &Seeder{
		Seeding: seeding,
		API:     api,
		Factory: NewFactory(seeding.Instruments, seeding.Difficulties)}
}

// Seed attempts to seed the database with the minimally required values and dummy data.
func (s *Seeder) Seed() {
	s.minimumSeed()
	s.dummySeed()
}

// minimumSeed when permitted, seeds the database with the minimally required values.
func (s *Seeder) minimumSeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableMinimumSeedingCommand) {
		log.Println("Did not seed the database with the minimally required values, as it was disabled.")
		return
	}
	s.seedAdmin()
	s.seedInstruments()
	s.seedDifficulties()
	s.seedSources()
	s.seedEndpoints()
}

// seedAdmin inserts an administrator user into the database.
func (s *Seeder) seedAdmin() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	if email != "" && password != "" {
		s.API.Users.InsertUser(email, password)
	} else {
		log.Println("Did not insert the initial admin account as no credentials were defined")
	}
}

// seedInstruments seeds the instruments table with the default instruments.
func (s *Seeder) seedInstruments() {
	s.API.Instruments.InsertInstruments(s.Seeding.Instruments...)
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.Difficulties.InsertDifficulties(s.Seeding.Difficulties...)
}

// seedSources seeds the sources from the config file.
func (s *Seeder) seedSources() {
	s.API.Sources.InsertSources(s.Seeding.Sources...)
}

// seedEndpoints seeds the endpoints from the config file.
func (s *Seeder) seedEndpoints() {
	s.API.Endpoints.InsertEndpoints(s.Seeding.Endpoints...)
}

// dummySeed when permitted, seeds the database with dummy data.
func (s *Seeder) dummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableDummySeedingCommand) {
		log.Println("Did not seed the database with dummy data, as it was disabled.")
		return
	}
	artists := s.Factory.CreateDummyArtists(s.Seeding.Dummies.Artists)
	// Insert the artists
	s.API.Artists.InsertArtists(artists...)
	for _, artist := range artists {
		// Insert the tracks of an artist
		s.API.Tracks.InsertTracks(artist.Tracks...)
		for _, track := range artist.Tracks {
			// Insert the tabs of a track
			s.API.Tabs.InsertTabs(track.Tabs...)
		}
	}
}

//TODO: Create seperate functions for seeding, each accepting arguments. Tracks do it for the provided art and tabs do it for the provided trcks.
//TODO: Create a seeder for the references table. One for art, one for trcks (the same as artist), and one for tabs.
//TODO: Put error messages somewhere, string format
//TODO: Better deal with errors and logging
//TODO: Add unit tests
//TODO: Add integration tests
//TODO: Create views?
