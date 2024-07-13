package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	api "github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"log"
)

// Creator helps with deleting data from the database
type Creator struct {
	API *api.API
}

// NewCreator creates a new instance of Creator
func NewCreator(API *api.API) *Creator {
	return &Creator{API: API}
}

// Create when permitted, creates tables in the database
func (c *Creator) Create() {
	if !logic.GetEnvValueAsBoolean(KeyDatabaseEnableCreatingCommand) {
		log.Fatalf("Did not create new tables for the database, as it was disabled")
	}
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
	c.CreateViews()
}

// CreateLookupTables creates the lookup tables.
func (c *Creator) CreateLookupTables() {
	c.API.Instruments.CreateInstrumentsTable()
	c.API.Difficulties.CreateDifficultiesTable()
	c.API.Sources.CreateSourcesTable()
	c.API.Endpoints.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *Creator) CreateStorageTables() {
	c.API.Artists.CreateArtistsTable()
	c.API.References.CreateReferencesTable()
	c.API.Sessions.CreateSessionsTable()
	c.API.Tabs.CreateTabsTable()
	c.API.Tracks.CreateTracksTable()
	c.API.Users.CreateUsersTable()
}

// CreateRelationLinkTables removes the relationship links between artists and tracks by creating and dropping the necessary tables.
func (c *Creator) CreateRelationLinkTables() {
	c.API.Artists.CreateArtistTrackTable()
	c.API.Tracks.CreateTrackTabTable()
}

// CreateViews creates the views.
func (c *Creator) CreateViews() {
	c.API.Views.CreateArtistsTracksTabsView()
	c.API.Views.CreateSourcesEndpointsView()
}
