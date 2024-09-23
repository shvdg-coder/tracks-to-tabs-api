package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ArtistTrackSchema represents schema operations related to 'artists to tracks' in the database.
type ArtistTrackSchema interface {
	CreateArtistTrackTable()
	DropArtistTrackTable()
}

// ArtistTrackSvc is for managing 'artists to tracks' links.
type ArtistTrackSvc struct {
	logic.DbOps
}

// NewArtistTrackSvc creates a new instance of the ArtistTrackSvc struct.
func NewArtistTrackSvc(database logic.DbOps) ArtistTrackSchema {
	return &ArtistTrackSvc{database}
}

// CreateArtistTrackTable creates an artist_track table if it doesn't already exist.
func (s *ArtistTrackSvc) CreateArtistTrackTable() {
	_, err := s.DB().Exec(queries.CreateArtistTrackTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropArtistTrackTable drops the artist_track table if it exists.
func (s *ArtistTrackSvc) DropArtistTrackTable() {
	_, err := s.DB().Exec(queries.DropArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
