package services

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"log"
)

// PurgeOperations represents operations for dropping tables.
type PurgeOperations interface {
	DropRelationLinkTables()
	DropStorageTables()
	DropLookupTables()
}

// PurgeService helps with deleting data from the database
type PurgeService struct {
	TableService *TableService
}

// NewPurgeService creates a new instance of PurgeService
func NewPurgeService(service *TableService) *PurgeService {
	return &PurgeService{TableService: service}
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
	p.TableService.DropArtistTrackTable()
	p.TableService.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *PurgeService) DropStorageTables() {
	p.TableService.DropArtistsTable()
	p.TableService.DropEndpointsTable()
	p.TableService.DropReferencesTable()
	p.TableService.DropSessionsTable()
	p.TableService.DropTracksTable()
	p.TableService.DropTabsTable()
	p.TableService.DropInstrumentsTable()
	p.TableService.DropUsersTable()
}

// DropLookupTables drops the lookup tables.
func (p *PurgeService) DropLookupTables() {
	p.TableService.DropInstrumentsTable()
	p.TableService.DropDifficultiesTable()
	p.TableService.DropSourcesTable()
}
