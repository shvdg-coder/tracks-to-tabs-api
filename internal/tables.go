package internal

import (
	"github.com/shvdg-dev/base-logic/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/internal/artists"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/internal/artists/artisttrack"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/internal/difficulties"
	end "github.com/shvdg-dev/tunes-to-tabs-api/internal/endpoints"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/internal/instruments"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/internal/references"
	sess "github.com/shvdg-dev/tunes-to-tabs-api/internal/sessions"
	src "github.com/shvdg-dev/tunes-to-tabs-api/internal/sources"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/internal/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/internal/tracks"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/internal/tracks/tracktab"
	usrs "github.com/shvdg-dev/tunes-to-tabs-api/internal/users"
)

// CreateTableOperations represents operations related to creating tables in the database.
type CreateTableOperations interface {
	CreateArtistsTable()
	CreateTracksTable()
	CreateArtistTrackTable()
	CreateTabsTable()
	CreateTrackTabTable()
	CreateUsersTable()
	CreateSessionsTable()
	CreateReferencesTable()
	CreateDifficultiesTable()
	CreateInstrumentsTable()
	CreateSourcesTable()
	CreateEndpointsTable()
}

// DropTableOperations represents operations related to dropping tables in the database.
type DropTableOperations interface {
	DropArtistsTable()
	DropTracksTable()
	DropArtistTrackTable()
	DropTabsTable()
	DropTrackTabTable()
	DropUsersTable()
	DropSessionsTable()
	DropReferencesTable()
	DropDifficultiesTable()
	DropInstrumentsTable()
	DropSourcesTable()
	DropEndpointsTable()
}

// TableOperations represents all operations related to the tables in the database.
type TableOperations interface {
	CreateTableOperations
	DropTableOperations
}

// TableService is responsible for maintaining the tables for the database.
type TableService struct {
	artistsOps      art.SetupOperations
	artistTracksOps arttrk.SetupOperations
	difficultiesOps diff.SetupOperations
	endpointsOps    end.SetupOperations
	instrumentsOps  inst.SetupOperations
	referencesOps   ref.SetupOperations
	sessionsOps     sess.SetupOperations
	sourcesOps      src.SetupOperations
	tabsOps         tbs.SetupOperations
	trackOps        trk.SetupOperations
	trackTabsOps    trktab.SetupOperations
	usersOps        usrs.SetupOperations
}

// NewTableService create a new instance of TableService.
func NewTableService(database pkg.DbOperations) TableOperations {
	return &TableService{
		artistsOps:      art.NewSetupService(database),
		trackOps:        trk.NewSetupService(database),
		artistTracksOps: arttrk.NewSetupService(database),
		tabsOps:         tbs.NewSetupService(database),
		trackTabsOps:    trktab.NewSetupService(database),
		usersOps:        usrs.NewSetupService(database),
		sessionsOps:     sess.NewSetupService(database),
		referencesOps:   ref.NewSetupService(database),
		difficultiesOps: diff.NewSetupService(database),
		instrumentsOps:  inst.NewSetupService(database),
		sourcesOps:      src.NewSetupService(database),
		endpointsOps:    end.NewSetupService(database),
	}
}

// CreateArtistsTable creates the artists table
func (ts *TableService) CreateArtistsTable() {
	ts.artistsOps.CreateArtistsTable()
}

// CreateTracksTable creates the tracks table
func (ts *TableService) CreateTracksTable() {
	ts.trackOps.CreateTracksTable()
}

// CreateArtistTrackTable creates the artist-track relationship table
func (ts *TableService) CreateArtistTrackTable() {
	ts.artistTracksOps.CreateArtistTrackTable()
}

// CreateTabsTable creates the tabs table
func (ts *TableService) CreateTabsTable() {
	ts.tabsOps.CreateTabsTable()
}

// CreateTrackTabTable creates the track-tab relationship table
func (ts *TableService) CreateTrackTabTable() {
	ts.trackTabsOps.CreateTrackTabTable()
}

// CreateUsersTable creates the users table
func (ts *TableService) CreateUsersTable() {
	ts.usersOps.CreateUsersTable()
}

// CreateSessionsTable creates the sessions table
func (ts *TableService) CreateSessionsTable() {
	ts.sessionsOps.CreateSessionsTable()
}

// CreateReferencesTable creates the references table
func (ts *TableService) CreateReferencesTable() {
	ts.referencesOps.CreateReferencesTable()
}

// CreateDifficultiesTable creates the difficulties table
func (ts *TableService) CreateDifficultiesTable() {
	ts.difficultiesOps.CreateDifficultiesTable()
}

// CreateInstrumentsTable creates the instruments table
func (ts *TableService) CreateInstrumentsTable() {
	ts.instrumentsOps.CreateInstrumentsTable()
}

// CreateSourcesTable creates the sources table
func (ts *TableService) CreateSourcesTable() {
	ts.sourcesOps.CreateSourcesTable()
}

// CreateEndpointsTable creates the endpoints table
func (ts *TableService) CreateEndpointsTable() {
	ts.endpointsOps.CreateEndpointsTable()
}

// DropArtistsTable drops the artists table
func (ts *TableService) DropArtistsTable() {
	ts.artistsOps.DropArtistsTable()
}

// DropTracksTable drops the tracks table
func (ts *TableService) DropTracksTable() {
	ts.trackOps.DropTracksTable()
}

// DropArtistTrackTable drops the artist-track relationship table
func (ts *TableService) DropArtistTrackTable() {
	ts.artistTracksOps.DropArtistTrackTable()
}

// DropTabsTable drops the tabs table
func (ts *TableService) DropTabsTable() {
	ts.tabsOps.DropTabsTable()
}

// DropTrackTabTable drops the track-tab relationship table
func (ts *TableService) DropTrackTabTable() {
	ts.trackTabsOps.DropTrackTabTable()
}

// DropUsersTable drops the users table
func (ts *TableService) DropUsersTable() {
	ts.usersOps.DropUsersTable()
}

// DropSessionsTable drops the sessions table
func (ts *TableService) DropSessionsTable() {
	ts.sessionsOps.DropSessionsTable()
}

// DropReferencesTable drops the references table
func (ts *TableService) DropReferencesTable() {
	ts.referencesOps.DropReferencesTable()
}

// DropDifficultiesTable drops the difficulties table
func (ts *TableService) DropDifficultiesTable() {
	ts.difficultiesOps.DropDifficultiesTable()
}

// DropInstrumentsTable drops the instruments table
func (ts *TableService) DropInstrumentsTable() {
	ts.instrumentsOps.DropInstrumentsTable()
}

// DropSourcesTable drops the sources table
func (ts *TableService) DropSourcesTable() {
	ts.sourcesOps.DropSourcesTable()
}

// DropEndpointsTable drops the endpoints table
func (ts *TableService) DropEndpointsTable() {
	ts.endpointsOps.DropEndpointsTable()
}
