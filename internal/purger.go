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
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnablePurgingCommand) {
		log.Fatalf("It is not allowed to purge the database")
	}
	p.DropViews()
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
	p.API.Endpoints.DropEndpointsTable()
	p.API.References.DropReferencesTable()
	p.API.Sessions.DropSessionsTable()
	p.API.Tracks.DropTracksTable()
	p.API.Tabs.DropTabsTable()
	p.API.Instruments.DropInstrumentsTable()
	p.API.Users.DropUsersTable()
}

// DropLookupTables drops the lookup tables.
func (p *Purger) DropLookupTables() {
	p.API.Instruments.DropInstrumentsTable()
	p.API.Difficulties.DropDifficultiesTable()
	p.API.Sources.DropSourcesTable()
}

// DropViews drops the views.
func (p *Purger) DropViews() {
	p.API.Tabs.DropTabsView()
	p.API.Endpoints.DropSourcesEndpointsView()
}
