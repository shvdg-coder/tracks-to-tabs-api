package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"log"
)

// Seeder helps with deleting data from the database
type Seeder struct {
	API *api.API
}

// NewSeeder creates a new instance of Seeder
func NewSeeder(API *api.API) *Seeder {
	return &Seeder{API: API}
}

// SeedTables attempts to seed the database with the minimally required values and dummy data.
func (s *Seeder) SeedTables() {
	s.MinimumSeed()
	s.DummySeed()
}

// MinimumSeed when permitted, seeds the database with the minimally required values.
func (s *Seeder) MinimumSeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowMinimumSeedingCommand) {
		log.Println("It is not allowed to seed the database with the minimally required values.")
		return
	}
	s.seedAdmin()
	s.seedInstruments()
	s.seedDifficulties()
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
	s.API.Instruments.InsertInstruments(
		inst.NewInstrument(InstrumentElectricGuitar),
		inst.NewInstrument(InstrumentAcousticGuitar),
		inst.NewInstrument(InstrumentBassGuitar),
		inst.NewInstrument(InstrumentDrums))
}

// seedDifficulties seeds the difficulties table with the default difficulties.
func (s *Seeder) seedDifficulties() {
	s.API.Difficulties.InsertDifficulties(
		diff.NewDifficulty(DifficultyEasy),
		diff.NewDifficulty(DifficultyIntermediate),
		diff.NewDifficulty(DifficultyHard),
		diff.NewDifficulty(DifficultyExpert))
}

// DummySeed when permitted, seeds the database with dummy data.
func (s *Seeder) DummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowDummySeedingCommand) {
		log.Println("It is not allowed to seed the database with dummy data.")
		return
	}
}

func (s *Seeder) seedDummyArtists() {
	//TODO:
}
