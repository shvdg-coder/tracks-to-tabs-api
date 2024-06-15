package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Purger helps with deleting data from the database
type Purger struct {
	API *api.API
}

// NewPurger creates a new instance of Purger
func NewPurger(API *api.API) *Purger {
	return &Purger{API: API}
}

// DropTables when permitted, drops the tables in the database
func (p *Purger) DropTables() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseAllowPurgingCommand) {
		log.Fatalf("It is not allowed to purge the database")
	}

	// Relation mappings
	p.API.Artists.DropArtistTrackTable()
	p.API.Tracks.DropTrackTabTable()

	// Data storage
	p.API.Artists.DropArtistsTable()
	p.API.IdReferences.DropIdReferencesTable()
	p.API.Instruments.DropInstrumentsTable()
	p.API.Sessions.DropSessionsTable()
	p.API.Tabs.DropTabsTable()
	p.API.Tracks.DropTracksTable()
	p.API.Users.DropUsersTable()

	// Lookup table
	p.API.Instruments.DropInstrumentsTable()
	p.API.Difficulties.DropDifficultiesTable()
}
