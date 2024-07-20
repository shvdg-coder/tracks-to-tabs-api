package services

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"log"
)

// PurgeService helps with deleting data from the database
type PurgeService struct {
	Service *TableService
}

// NewPurgeService creates a new instance of PurgeService
func NewPurgeService(service *TableService) *PurgeService {
	return &PurgeService{Service: service}
}

// Purge when permitted, drops the tables in the database
func (p *PurgeService) Purge() {
	if !logic.GetEnvValueAsBoolean(internal.KeyDatabaseEnablePurgingCommand) {
		log.Fatalf("It is not allowed to purge the database")
	}
	p.DropRelationLinkTables()
	p.DropStorageTables()
	p.DropLookupTables()
}

// DropRelationLinkTables drops the tables that hold relation links.
func (p *PurgeService) DropRelationLinkTables() {
	p.Service.DropArtistTrackTable()
	p.Service.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *PurgeService) DropStorageTables() {
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
func (p *PurgeService) DropLookupTables() {
	p.Service.DropInstrumentsTable()
	p.Service.DropDifficultiesTable()
	p.Service.DropSourcesTable()
}
