package artists

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to artists in the database.
type SetupOperations interface {
	CreateArtistsTable()
	DropArtistsTable()
}

// SetupService is for managing artists.
type SetupService struct {
	*logic.DatabaseManager
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database *logic.DatabaseManager) SetupOperations {
	return &SetupService{database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (s *SetupService) CreateArtistsTable() {
	_, err := s.DB.Exec(CreateArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artists' table")
	}
}

// DropArtistsTable drops the artists table if it exists.
func (s *SetupService) DropArtistsTable() {
	_, err := s.DB.Exec(DropArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artists' table")
	}
}
