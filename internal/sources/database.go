package sources

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to sources in the database.
type SetupOperations interface {
	CreateSourcesTable()
	DropSourcesTable()
}

// SetupService is for setting up the sources table in the database.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{DbOperations: database}
}

// CreateSourcesTable creates a sources table if it doesn't already exist.
func (s *SetupService) CreateSourcesTable() {
	_, err := s.Exec(CreateSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' table.")
	}
}

// DropSourcesTable drops the sources table if it exists.
func (s *SetupService) DropSourcesTable() {
	_, err := s.Exec(DropSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' table.")
	}
}
