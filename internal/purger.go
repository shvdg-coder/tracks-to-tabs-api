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
	p.DropRelationLinkTables()
	p.DropStorageTables()
	p.DropLookupTables()
}

// DropRelationLinkTables drops the tables that hold relation links.
func (p *Purger) DropRelationLinkTables() {
	p.API.Artists.DropArtistTrackTable()
	p.API.Tracks.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *Purger) DropStorageTables() {
	p.API.Artists.DropArtistsTable()
	p.API.References.DropResourcesTable()
	p.API.Instruments.DropInstrumentsTable()
	p.API.Sessions.DropSessionsTable()
	p.API.Tabs.DropTabsTable()
	p.API.Tracks.DropTracksTable()
	p.API.Users.DropUsersTable()
}

// DropLookupTables drops the lookup tables.
func (p *Purger) DropLookupTables() {
	p.API.Instruments.DropInstrumentsTable()
	p.API.Difficulties.DropDifficultiesTable()
}
