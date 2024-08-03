package internal

// CreateOperations represents operations for creating tables.
type CreateOperations interface {
	CreateAll()
	CreateLookupTables()
	CreateStorageTables()
	CreateRelationLinkTables()
}

// CreateService helps with creating tables for the database
type CreateService struct {
	TableOps TableOperations
}

// NewCreateService creates a new instance of CreateService
func NewCreateService(tables TableOperations) CreateOperations {
	return &CreateService{TableOps: tables}
}

// CreateAll when permitted, creates tables in the database
func (c *CreateService) CreateAll() {
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *CreateService) CreateLookupTables() {
	c.TableOps.CreateInstrumentsTable()
	c.TableOps.CreateDifficultiesTable()
	c.TableOps.CreateSourcesTable()
	c.TableOps.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *CreateService) CreateStorageTables() {
	c.TableOps.CreateArtistsTable()
	c.TableOps.CreateReferencesTable()
	c.TableOps.CreateSessionsTable()
	c.TableOps.CreateTabsTable()
	c.TableOps.CreateTracksTable()
	c.TableOps.CreateUsersTable()
}

// CreateRelationLinkTables creates tables for establishing links.
func (c *CreateService) CreateRelationLinkTables() {
	c.TableOps.CreateArtistTrackTable()
	c.TableOps.CreateTrackTabTable()
}
