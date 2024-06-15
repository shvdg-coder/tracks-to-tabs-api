package tabs

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"

	"log"
)

// API is for managing tabs.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateTabsTable creates a tabs table if it doesn't already exist.
func (a *API) CreateTabsTable() {
	_, err := a.Database.DB.Exec(createTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tabs' table.")
	}
}

// DropTabsTable drops the tabs table if it exists.
func (a *API) DropTabsTable() {
	_, err := a.Database.DB.Exec(dropTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tabs' table.")
	}
}

// InsertTab inserts a new tab in the tabs table.
func (a *API) InsertTab(instrumentId int, difficultyId int, tuningId int, description string) {
	_, err := a.Database.DB.Exec(insertTabQuery, instrumentId, difficultyId, tuningId, description)
	if err != nil {
		log.Printf("Failed inserting tab with Following IDs '%d', '%d', '%d' & Description: '%s': %s", instrumentId, difficultyId, tuningId, description, err.Error())
	} else {
		log.Printf("Successfully inserted tab with Following IDs '%d', '%d', '%d' & Description: '%s'", instrumentId, difficultyId, tuningId, description)
	}
}
