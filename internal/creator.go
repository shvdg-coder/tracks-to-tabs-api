package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal/services"
	"log"
)

// Creator helps with deleting data from the database
type Creator struct {
	Service *services.TableService
}

// NewCreator creates a new instance of Creator
func NewCreator(service *services.TableService) *Creator {
	return &Creator{Service: service}
}

// Create when permitted, creates tables in the database
func (c *Creator) Create() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableCreatingCommand) {
		log.Fatalf("Did not create new tables for the database, as it was disabled")
	}
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *Creator) CreateLookupTables() {
	c.Service.CreateInstrumentsTable()
	c.Service.CreateDifficultiesTable()
	c.Service.CreateSourcesTable()
	c.Service.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *Creator) CreateStorageTables() {
	c.Service.CreateArtistsTable()
	c.Service.CreateReferencesTable()
	c.Service.CreateSessionsTable()
	c.Service.CreateTabsTable()
	c.Service.CreateTracksTable()
	c.Service.CreateUsersTable()
}

// CreateRelationLinkTables removes the relationship links between artists and tracks by creating and dropping the necessary tables.
func (c *Creator) CreateRelationLinkTables() {
	c.Service.CreateArtistTrackTable()
	c.Service.CreateTrackTabTable()
}
