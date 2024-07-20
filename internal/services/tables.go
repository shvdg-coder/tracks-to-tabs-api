package services

import (
	"github.com/shvdg-dev/base-logic/pkg"
	quer "github.com/shvdg-dev/tunes-to-tabs-api/internal/queries"
	"log"
)

// TableService is responsible for maintaining the tables for the database.
type TableService struct {
	Database *pkg.DatabaseManager
}

// NewTableService create a new instance of TableService.
func NewTableService(database *pkg.DatabaseManager) *TableService {
	return &TableService{Database: database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (t *TableService) CreateArtistsTable() {
	_, err := t.Database.DB.Exec(quer.CreateArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artists' table")
	}
}

// DropArtistsTable drops the artists table if it exists.
func (t *TableService) DropArtistsTable() {
	_, err := t.Database.DB.Exec(quer.DropArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artists' table")
	}
}

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (t *TableService) CreateTracksTable() {
	_, err := t.Database.DB.Exec(quer.CreateTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tracks' table")
	}
}

// DropTracksTable drops the tracks table if it exists.
func (t *TableService) DropTracksTable() {
	_, err := t.Database.DB.Exec(quer.DropTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tracks' table")
	}
}

// CreateArtistTrackTable creates an artist_track table if it doesn't already exist.
func (t *TableService) CreateArtistTrackTable() {
	_, err := t.Database.DB.Exec(quer.CreateArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artist_track' table")
	}
}

// DropArtistTrackTable drops the artist_track table if it exists.
func (t *TableService) DropArtistTrackTable() {
	_, err := t.Database.DB.Exec(quer.DropArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artist_track' table")
	}
}

// CreateDifficultiesTable creates a difficulties table.
func (t *TableService) CreateDifficultiesTable() {
	_, err := t.Database.DB.Exec(quer.CreateDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'difficulties' table.")
	}
}

// DropDifficultiesTable drops the difficulties table.
func (t *TableService) DropDifficultiesTable() {
	_, err := t.Database.DB.Exec(quer.DropDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'difficulties' table.")
	}
}

// CreateEndpointsTable creates the endpoints table if it doesn't already exist.
func (t *TableService) CreateEndpointsTable() {
	_, err := t.Database.DB.Exec(quer.CreateEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'endpoints' table")
	}
}

// DropEndpointsTable drops the endpoints table if it exists.
func (t *TableService) DropEndpointsTable() {
	_, err := t.Database.DB.Exec(quer.DropEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'endpoints' table")
	}
}

// CreateInstrumentsTable creates an instruments table if it doesn't already exist.
func (t *TableService) CreateInstrumentsTable() {
	_, err := t.Database.DB.Exec(quer.CreateInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'instruments' table.")
	}
}

// DropInstrumentsTable drops the instruments table if it exists.
func (t *TableService) DropInstrumentsTable() {
	_, err := t.Database.DB.Exec(quer.DropInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'instruments' table.")
	}
}

// CreateReferencesTable creates the references table if it doesn't already exist.
func (t *TableService) CreateReferencesTable() {
	_, err := t.Database.DB.Exec(quer.CreateReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'references' table")
	}
}

// DropReferencesTable drops the references table if it exists.
func (t *TableService) DropReferencesTable() {
	_, err := t.Database.DB.Exec(quer.DropReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'references' table")
	}
}

// CreateSessionsTable creates the sessions table in the database and adds an expiry index.
func (t *TableService) CreateSessionsTable() {
	_, err := t.Database.DB.Exec(quer.CreateSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = t.Database.DB.Exec(quer.CreateSessionExpiryIndexQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sessions' table")
	}
}

// DropSessionsTable drops the sessions table if it exists.
func (t *TableService) DropSessionsTable() {
	_, err := t.Database.DB.Exec(quer.DropSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sessions' table")
	}
}

// CreateSourcesTable creates a sources table if it doesn't already exist.
func (t *TableService) CreateSourcesTable() {
	_, err := t.Database.DB.Exec(quer.CreateSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' table.")
	}
}

// DropSourcesTable drops the sources table if it exists.
func (t *TableService) DropSourcesTable() {
	_, err := t.Database.DB.Exec(quer.DropSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' table.")
	}
}

// CreateTabsTable creates a tabs table if it doesn't already exist.
func (t *TableService) CreateTabsTable() {
	_, err := t.Database.DB.Exec(quer.CreateTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tabs' table.")
	}
}

// DropTabsTable drops the tabs table if it exists.
func (t *TableService) DropTabsTable() {
	_, err := t.Database.DB.Exec(quer.DropTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tabs' table.")
	}
}

// CreateTrackTabTable creates a track_tab table if it doesn't already exist.
func (t *TableService) CreateTrackTabTable() {
	_, err := t.Database.DB.Exec(quer.CreateTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'track_tab' table")
	}
}

// DropTrackTabTable drops the track_tab table if it exists.
func (t *TableService) DropTrackTabTable() {
	_, err := t.Database.DB.Exec(quer.DropTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'track_tab' table")
	}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (t *TableService) CreateUsersTable() {
	_, err := t.Database.DB.Exec(quer.CreateUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'users' table")
	}
}

// DropUsersTable drops the users table if it exists.
func (t *TableService) DropUsersTable() {
	_, err := t.Database.DB.Exec(quer.DropUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'users' table")
	}
}
