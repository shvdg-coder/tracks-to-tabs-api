package difficulties

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to difficulties in the database.
type SetupOperations interface {
	CreateDifficultiesTable()
	DropDifficultiesTable()
}

// SetupService is for setting up difficulties in the database.
type SetupService struct {
	*logic.DatabaseManager
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database *logic.DatabaseManager) SetupOperations {
	return &SetupService{DatabaseManager: database}
}

// CreateDifficultiesTable creates a difficulties table.
func (s *SetupService) CreateDifficultiesTable() {
	_, err := s.DB.Exec(CreateDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'difficulties' table.")
	}
}

// DropDifficultiesTable drops the difficulties table.
func (s *SetupService) DropDifficultiesTable() {
	_, err := s.DB.Exec(DropDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'difficulties' table.")
	}
}
