package services

// DropOperations represents operations for dropping tables.
type DropOperations interface {
	DropAll()
	DropRelationLinkTables()
	DropStorageTables()
	DropLookupTables()
}

// DropService helps with deleting data from the database
type DropService struct {
	TableService *TableService
}

// NewDropService creates a new instance of DropService
func NewDropService(service *TableService) *DropService {
	return &DropService{TableService: service}
}

// DropAll when permitted, drops the tables in the database
func (p *DropService) DropAll() {
	p.DropRelationLinkTables()
	p.DropStorageTables()
	p.DropLookupTables()
}

// DropRelationLinkTables drops the tables that hold relation links.
func (p *DropService) DropRelationLinkTables() {
	p.TableService.DropArtistTrackTable()
	p.TableService.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *DropService) DropStorageTables() {
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
func (p *DropService) DropLookupTables() {
	p.TableService.DropInstrumentsTable()
	p.TableService.DropDifficultiesTable()
	p.TableService.DropSourcesTable()
}
