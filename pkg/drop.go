package pkg

import logic "github.com/shvdg-dev/base-logic/pkg"

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
func NewDropAPI(database logic.DbOperations) DropOps {
	return &DropAPI{SvcOps: NewSvcManager(database)}
}

// DropAll when permitted, drops the tables in the database.
func (d *DropAPI) DropAll() {
	d.DropLookupTables()
	d.DropStorageTables()
	d.DropRelationLinkTables()
}

// DropLookupTables drops the lookup tables.
func (d *DropAPI) DropLookupTables() {
	d.DropInstrumentsTable()
	d.DropDifficultiesTable()
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
