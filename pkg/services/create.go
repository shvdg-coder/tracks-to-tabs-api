package services

// CreateOps represents operations for creating tables.
type CreateOps interface {
	CreateAll()
	CreateLookupTables()
	CreateStorageTables()
	CreateRelationLinkTables()
}

// CreateSvc helps with creating tables for the database
type CreateSvc struct {
	SvcOps
}

// NewCreateSvc creates a new instance of CreateSvc
func NewCreateSvc(svcManager SvcOps) CreateOps {
	return &CreateSvc{SvcOps: svcManager}
}

// CreateAll when permitted, creates tables in the database
func (c *CreateSvc) CreateAll() {
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *CreateSvc) CreateLookupTables() {
	c.CreateInstrumentsTable()
	c.CreateDifficultiesTable()
	c.CreateSourcesTable()
	c.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *CreateSvc) CreateStorageTables() {
	c.CreateArtistsTable()
	c.CreateReferencesTable()
	c.CreateTabsTable()
	c.CreateTracksTable()
	c.CreateUsersTable()
}

// CreateRelationLinkTables creates tables for establishing links.
func (c *CreateSvc) CreateRelationLinkTables() {
	c.CreateArtistTrackTable()
	c.CreateTrackTabTable()
}
