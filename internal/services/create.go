package services

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"log"
)

// CreateService helps with deleting data from the database
type CreateService struct {
	Service *TableService
}

// NewCreateService creates a new instance of CreateService
func NewCreateService(service *TableService) *CreateService {
	return &CreateService{Service: service}
}

// CreateAll when permitted, creates tables in the database
func (c *CreateService) CreateAll() {
	if !logic.GetEnvValueAsBoolean(internal.KeyDatabaseEnableCreatingCommand) {
		log.Fatalf("Did not create new tables for the database, as it was disabled")
	}
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *CreateService) CreateLookupTables() {
	c.Service.CreateInstrumentsTable()
	c.Service.CreateDifficultiesTable()
	c.Service.CreateSourcesTable()
	c.Service.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *CreateService) CreateStorageTables() {
	c.Service.CreateArtistsTable()
	c.Service.CreateReferencesTable()
	c.Service.CreateSessionsTable()
	c.Service.CreateTabsTable()
	c.Service.CreateTracksTable()
	c.Service.CreateUsersTable()
}

// CreateRelationLinkTables removes the relationship links between artists and tracks by creating and dropping the necessary tables.
func (c *CreateService) CreateRelationLinkTables() {
	c.Service.CreateArtistTrackTable()
	c.Service.CreateTrackTabTable()
}
