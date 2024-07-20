package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal/services"
	"log"
)

// Purger helps with deleting data from the database
type Purger struct {
	Service *services.TableService
}

// NewPurger creates a new instance of Purger
func NewPurger(service *services.TableService) *Purger {
	return &Purger{Service: service}
}

// Purge when permitted, drops the tables in the database
func (p *Purger) Purge() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnablePurgingCommand) {
		log.Fatalf("It is not allowed to purge the database")
	}
	p.DropRelationLinkTables()
	p.DropStorageTables()
	p.DropLookupTables()
}

// DropRelationLinkTables drops the tables that hold relation links.
func (p *Purger) DropRelationLinkTables() {
	p.Service.DropArtistTrackTable()
	p.Service.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *Purger) DropStorageTables() {
	p.Service.DropArtistsTable()
	p.Service.DropEndpointsTable()
	p.Service.DropReferencesTable()
	p.Service.DropSessionsTable()
	p.Service.DropTracksTable()
	p.Service.DropTabsTable()
	p.Service.DropInstrumentsTable()
	p.Service.DropUsersTable()
}

// DropLookupTables drops the lookup tables.
func (p *Purger) DropLookupTables() {
	p.Service.DropInstrumentsTable()
	p.Service.DropDifficultiesTable()
	p.Service.DropSourcesTable()
}
