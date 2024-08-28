package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ArtistSchema represents schema operations related to artists in the database.
type ArtistSchema interface {
	CreateArtistsTable()
	DropArtistsTable()
}

// ArtistSvc is for managing artists.
type ArtistSvc struct {
	logic.DbOperations
}

// NewArtistSvc creates a new instance of the ArtistSvc struct.
func NewArtistSvc(database logic.DbOperations) ArtistSchema {
	return &ArtistSvc{database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (s *ArtistSvc) CreateArtistsTable() {
	_, err := s.Exec(queries.CreateArtistsTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropArtistsTable drops the artists table if it exists.
func (s *ArtistSvc) DropArtistsTable() {
	_, err := s.Exec(queries.DropArtistsTable)
	if err != nil {
		log.Fatal(err)
	}
}
