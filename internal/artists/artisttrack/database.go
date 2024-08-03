package artisttrack

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// SetupOperations represents setup operations related to 'artists to tracks'.
type SetupOperations interface {
	CreateArtistTrackTable()
	DropArtistTrackTable()
}

// SetupService is for managing 'artists to tracks' links.
type SetupService struct {
	logic.DbOperations
}

// NewSetupService creates a new instance of the SetupService struct.
func NewSetupService(database logic.DbOperations) SetupOperations {
	return &SetupService{DbOperations: database}
}

// CreateArtistTrackTable creates an artist_track table if it doesn't already exist.
func (s *SetupService) CreateArtistTrackTable() {
	_, err := s.Exec(CreateArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artist_track' table")
	}
}

// DropArtistTrackTable drops the artist_track table if it exists.
func (s *SetupService) DropArtistTrackTable() {
	_, err := s.Exec(DropArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artist_track' table")
	}
}
