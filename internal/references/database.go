package references

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to references in the database.
type SetupOperations interface {
	CreateReferencesTable()
	DropReferencesTable()
}

// SetupService is for setting up references table in the database.
type SetupService struct {
	*logic.DatabaseManager
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database *logic.DatabaseManager) SetupOperations {
	return &SetupService{DatabaseManager: database}
}

// CreateReferencesTable creates the references table if it doesn't already exist.
func (s *SetupService) CreateReferencesTable() {
	_, err := s.DB.Exec(CreateReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'references' table")
	}
}

// DropReferencesTable drops the references table if it exists.
func (s *SetupService) DropReferencesTable() {
	_, err := s.DB.Exec(DropReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'references' table")
	}
}
