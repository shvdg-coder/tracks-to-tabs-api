package endpoints

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to endpoints in the database.
type SetupOperations interface {
	CreateEndpointsTable()
	DropEndpointsTable()
}

// SetupService is for setting up endpoints tables in the database.
type SetupService struct {
	*logic.DatabaseManager
}

// NewSetupService creates a new instance of SetupService.
func NewSetupService(database *logic.DatabaseManager) SetupOperations {
	return &SetupService{DatabaseManager: database}
}

// CreateEndpointsTable creates the endpoints table if it doesn't already exist.
func (t *SetupService) CreateEndpointsTable() {
	_, err := t.DB.Exec(CreateEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'endpoints' table")
	}
}

// DropEndpointsTable drops the endpoints table if it exists.
func (t *SetupService) DropEndpointsTable() {
	_, err := t.DB.Exec(DropEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'endpoints' table")
	}
}
