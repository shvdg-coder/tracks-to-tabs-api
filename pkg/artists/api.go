package artists

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"

	"log"
)

// API is for managing users.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (a *API) CreateArtistsTable() {
	_, err := a.Database.DB.Exec(createArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the Artists table")
	}
}

// DropArtistsTable drops the artists table if it exists.
func (a *API) DropArtistsTable() {
	_, err := a.Database.DB.Exec(dropArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the Artists table")
	}
}

// InsertArtist inserts an artist into the artists table.
func (a *API) InsertArtist(name string) {
	_, err := a.Database.DB.Exec(insertArtistQuery, name)
	if err != nil {
		log.Printf("Failed inserting user with name '%s': %s", name, err.Error())
	} else {
		log.Printf("Successfully inserted artist '%s' into the Artists table", name)
	}
}
