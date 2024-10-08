package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ArtistSchema represents schema operations related to artists in the database.
type ArtistSchema interface {
	CreateArtistsTable()
	DropArtistsTable()
}

// ArtistSvc is for managing artists.
type ArtistSvc struct {
	logic.DbOps
}

// NewArtistSvc creates a new instance of the ArtistSvc struct.
func NewArtistSvc(database logic.DbOps) ArtistSchema {
	return &ArtistSvc{database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (s *ArtistSvc) CreateArtistsTable() {
	_, err := s.DB().Exec(queries.CreateArtistsTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropArtistsTable drops the artists table if it exists.
func (s *ArtistSvc) DropArtistsTable() {
	_, err := s.DB().Exec(queries.DropArtistsTable)
	if err != nil {
		log.Fatal(err)
	}
}
