package services

// CreateOperations represents operations for creating tables.
type CreateOperations interface {
	CreateAll()
	CreateLookupTables()
	CreateStorageTables()
	CreateRelationLinkTables()
}

// CreateService helps with creating tables for the database
type CreateService struct {
	TableService *TableService
}

// NewCreateService creates a new instance of CreateService
func NewCreateService(service *TableService) *CreateService {
	return &CreateService{TableService: service}
}

// CreateAll when permitted, creates tables in the database
func (c *CreateService) CreateAll() {
	c.CreateLookupTables()
	c.CreateStorageTables()
	c.CreateRelationLinkTables()
}

// CreateLookupTables creates the lookup tables.
func (c *CreateService) CreateLookupTables() {
	c.TableService.CreateInstrumentsTable()
	c.TableService.CreateDifficultiesTable()
	c.TableService.CreateSourcesTable()
	c.TableService.CreateEndpointsTable()
}

// CreateStorageTables creates tables.
func (c *CreateService) CreateStorageTables() {
	c.TableService.CreateArtistsTable()
	c.TableService.CreateReferencesTable()
	c.TableService.CreateSessionsTable()
	c.TableService.CreateTabsTable()
	c.TableService.CreateTracksTable()
	c.TableService.CreateUsersTable()
}

// CreateRelationLinkTables creates tables for establishing links.
func (c *CreateService) CreateRelationLinkTables() {
	c.TableService.CreateArtistTrackTable()
	c.TableService.CreateTrackTabTable()
}
