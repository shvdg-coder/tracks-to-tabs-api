package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
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
	s.InsertAdmin()
}

// InsertAdmin inserts an administrator user into the database.
func (s *Seeder) InsertAdmin() {
	email := logic.GetEnvValueAsString(KeyAdminInitialEmail)
	password := logic.GetEnvValueAsString(KeyAdminInitialPassword)
	s.API.Users.InsertUser(email, password)
}

// DummySeed when permitted, seeds the database with dummy data.
func (s *Seeder) DummySeed() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowDummySeedingCommand) {
		log.Println("It is not allowed to seed the database with dummy data.")
		return
	}
}
