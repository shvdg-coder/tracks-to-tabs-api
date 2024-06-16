package difficulties

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing difficulties.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateDifficultiesTable creates a difficulties table.
func (a *API) CreateDifficultiesTable() {
	_, err := a.Database.DB.Exec(createDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'difficulties' table.")
	}
}

// DropDifficultiesTable drops the difficulties table.
func (a *API) DropDifficultiesTable() {
	_, err := a.Database.DB.Exec(dropDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'difficulties' table.")
	}
}

// InsertDifficulty inserts a new difficulty level.
func (a *API) InsertDifficulty(name string) {
	_, err := a.Database.DB.Exec(insertDifficultyQuery, name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with Title: '%s': %s", name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with Title: '%s'", name)
	}
}
