package internal

// DropOperations represents operations for dropping tables.
type DropOperations interface {
	DropAll()
	DropRelationLinkTables()
	DropStorageTables()
	DropLookupTables()
}

// DropService helps with deleting data from the database
type DropService struct {
	TableOps TableOperations
}

// NewDropService creates a new instance of DropService
func NewDropService(tableOps TableOperations) DropOperations {
	return &DropService{TableOps: tableOps}
}

// DropAll when permitted, drops the tables in the database
func (p *DropService) DropAll() {
	p.DropRelationLinkTables()
	p.DropStorageTables()
	p.DropLookupTables()
}

// DropRelationLinkTables drops the tables that hold relation links.
func (p *DropService) DropRelationLinkTables() {
	p.TableOps.DropArtistTrackTable()
	p.TableOps.DropTrackTabTable()
}

// DropStorageTables drops tables.
func (p *DropService) DropStorageTables() {
	p.TableOps.DropArtistsTable()
	p.TableOps.DropEndpointsTable()
	p.TableOps.DropReferencesTable()
	p.TableOps.DropSessionsTable()
	p.TableOps.DropTracksTable()
	p.TableOps.DropTabsTable()
	p.TableOps.DropInstrumentsTable()
	p.TableOps.DropUsersTable()
}

// DropLookupTables drops the lookup tables.
func (p *DropService) DropLookupTables() {
	p.TableOps.DropInstrumentsTable()
	p.TableOps.DropDifficultiesTable()
	p.TableOps.DropSourcesTable()
}
