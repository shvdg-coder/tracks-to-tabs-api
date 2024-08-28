package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ArtistTrackSchema represents schema operations related to 'artists to tracks' in the database.
type ArtistTrackSchema interface {
	CreateArtistTrackTable()
	DropArtistTrackTable()
}

// ArtistTrackSvc is for managing 'artists to tracks' links.
type ArtistTrackSvc struct {
	logic.DbOperations
}

// NewArtistTrackSvc creates a new instance of the ArtistTrackSvc struct.
func NewArtistTrackSvc(database logic.DbOperations) ArtistTrackSchema {
	return &ArtistTrackSvc{database}
}

// CreateArtistTrackTable creates an artist_track table if it doesn't already exist.
func (s *ArtistTrackSvc) CreateArtistTrackTable() {
	_, err := s.Exec(queries.CreateArtistTrackTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropArtistTrackTable drops the artist_track table if it exists.
func (s *ArtistTrackSvc) DropArtistTrackTable() {
	_, err := s.Exec(queries.DropArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
