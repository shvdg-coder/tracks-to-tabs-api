package users

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to users in the database.
type SetupOperations interface {
	CreateUsersTable()
	DropUsersTable()
}

// SetupService is for setting up the users table.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{database}
}

// CreateUsersTable creates a users table if it doesn't already exist.
func (s *SetupService) CreateUsersTable() {
	_, err := s.Exec(CreateUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'users' table")
	}
}

// DropUsersTable drops the users table if it exists.
func (s *SetupService) DropUsersTable() {
	_, err := s.Exec(DropUsersTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'users' table")
	}
}
