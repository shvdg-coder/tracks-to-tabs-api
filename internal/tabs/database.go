package tabs

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to tabs in the database.
type SetupOperations interface {
	CreateTabsTable()
	DropTabsTable()
}

// SetupService is for setting up the tabs table in the database.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{DbOperations: database}
}

// CreateTabsTable creates a tabs table if it doesn't already exist.
func (s *SetupService) CreateTabsTable() {
	_, err := s.Exec(CreateTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tabs' table.")
	}
}

// DropTabsTable drops the tabs table if it exists.
func (s *SetupService) DropTabsTable() {
	_, err := s.Exec(DropTabsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tabs' table.")
	}
}
