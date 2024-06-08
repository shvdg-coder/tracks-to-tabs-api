package sessions

import (
	"github.com/shvdg-dev/base-pkg/database"
	"log"
)

// API is for managing users.
type API struct {
	Database *database.Manager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *database.Manager) *API {
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
	}
}
