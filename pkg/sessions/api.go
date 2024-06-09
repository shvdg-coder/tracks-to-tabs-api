package sessions

import (
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

// CreateSessionsTable creates the sessions table in the database and adds an expiry index.
func (a *API) CreateSessionsTable() {
	_, err := a.Database.DB.Exec(createSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = a.Database.DB.Exec(createSessionExpiryIndexQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the Sessions table")
	}
}

// DropSessionsTable drops the sessions table if it exists.
func (a *API) DropSessionsTable() {
	_, err := a.Database.DB.Exec(dropSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the Sessions table")
	}
}
