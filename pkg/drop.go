package pkg

// DropOps represents operations for dropping tables.
type DropOps interface {
	DropAll()
	DropLookupTables()
	DropStorageTables()
	DropRelationLinkTables()
}

// DropAPI helps with deleting data from the database.
type DropAPI struct {
	SvcOps
}

// NewDropAPI creates a new instance of DropAPI.
func NewDropAPI(svcManager SvcOps) DropOps {
	return &DropAPI{svcManager}
}

// DropAll when permitted, drops the tables in the database.
func (d *DropAPI) DropAll() {
	d.DropRelationLinkTables()
	d.DropStorageTables()
	d.DropLookupTables()
}

// DropLookupTables drops the lookup tables.
func (d *DropAPI) DropLookupTables() {
	d.DropInstrumentsTable()
	d.DropDifficultiesTable()
	d.DropEndpointsTable()
	d.DropSourcesTable()
}

// DropStorageTables drops the storage tables.
func (d *DropAPI) DropStorageTables() {
	d.DropArtistsTable()
	d.DropReferencesTable()
	d.DropTabsTable()
	d.DropTracksTable()
	d.DropUsersTable()
}

// DropRelationLinkTables drops tables that establish links.
func (d *DropAPI) DropRelationLinkTables() {
	d.DropArtistTrackTable()
	d.DropTrackTabTable()
}
