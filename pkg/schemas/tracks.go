package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TrackSchema represents schema operations related to 'tracks' in the database.
type TrackSchema interface {
	CreateTracksTable()
	DropTracksTable()
}

// TrackSvc is for managing 'tracks' tables in the database.
type TrackSvc struct {
	logic.DbOps
}

// NewTrackSvc creates a new instance of the TrackSvc struct.
func NewTrackSvc(database logic.DbOps) TrackSchema {
	return &TrackSvc{database}
}

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (s *TrackSvc) CreateTracksTable() {
	_, err := s.DB().Exec(queries.CreateTracksTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropTracksTable drops the tracks table if it exists.
func (s *TrackSvc) DropTracksTable() {
	_, err := s.DB().Exec(queries.DropTracksTable)
	if err != nil {
		log.Fatal(err)
	}
}
