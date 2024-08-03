package sessions

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to users in the database.
type SetupOperations interface {
	CreateSessionsTable()
	DropSessionsTable()
}

// SetupService is for setting up the users table.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{database}
}

// CreateSessionsTable creates the sessions table in the database and adds an expiry index.
func (s *SetupService) CreateSessionsTable() {
	_, err := s.Exec(CreateSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Exec(CreateSessionExpiryIndexQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sessions' table")
	}
}

// DropSessionsTable drops the sessions table if it exists.
func (s *SetupService) DropSessionsTable() {
	_, err := s.Exec(DropSessionsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sessions' table")
	}
}
