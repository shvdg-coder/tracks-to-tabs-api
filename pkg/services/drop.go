package services

// DropOps represents operations for dropping tables.
type DropOps interface {
	DropAll()
	DropLookupTables()
	DropStorageTables()
	DropRelationLinkTables()
}

// DropSvc helps with deleting data from the database.
type DropSvc struct {
	SvcOps
}

// NewDropSvc creates a new instance of DropSvc.
func NewDropSvc(svcManager SvcOps) DropOps {
	return &DropSvc{svcManager}
}

// DropAll when permitted, drops the tables in the database.
func (d *DropSvc) DropAll() {
	d.DropRelationLinkTables()
	d.DropStorageTables()
	d.DropLookupTables()
}

// DropLookupTables drops the lookup tables.
func (d *DropSvc) DropLookupTables() {
	d.DropInstrumentsTable()
	d.DropDifficultiesTable()
	d.DropEndpointsTable()
	d.DropSourcesTable()
}

// DropStorageTables drops the storage tables.
func (d *DropSvc) DropStorageTables() {
	d.DropArtistsTable()
	d.DropReferencesTable()
	d.DropTabsTable()
	d.DropTracksTable()
	d.DropUsersTable()
}

// DropRelationLinkTables drops tables that establish links.
func (d *DropSvc) DropRelationLinkTables() {
	d.DropArtistTrackTable()
	d.DropTrackTabTable()
}
