package pkg

import logic "github.com/shvdg-dev/base-logic/pkg"

// CreateOps represents operations for creating tables.
type CreateOps interface {
	CreateAll()
	CreateLookupTables()
	CreateStorageTables()
	CreateRelationLinkTables()
}

// CreateAPI helps with creating tables for the database
type CreateAPI struct {
	SvcOps
}

// NewCreateAPI creates a new instance of CreateAPI
func NewCreateAPI(database logic.DbOperations) CreateOps {
	return &CreateAPI{SvcOps: NewSvcManager(database)}
}

// CreateAll when permitted, creates tables in the database
func (c *CreateAPI) CreateAll() {
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *CreateAPI) CreateLookupTables() {
	c.CreateInstrumentsTable()
	c.CreateDifficultiesTable()
	c.CreateSourcesTable()
	c.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *CreateAPI) CreateStorageTables() {
	c.CreateArtistsTable()
	c.CreateReferencesTable()
	c.CreateTabsTable()
	c.CreateTracksTable()
	c.CreateUsersTable()
}

// CreateRelationLinkTables creates tables for establishing links.
func (c *CreateAPI) CreateRelationLinkTables() {
	c.CreateArtistTrackTable()
	c.CreateTrackTabTable()
}
