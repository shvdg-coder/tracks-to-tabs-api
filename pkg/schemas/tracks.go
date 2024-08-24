package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// TrackSchema represents schema operations related to 'tracks' in the database.
type TrackSchema interface {
	CreateTracksTable()
	DropTracksTable()
}

// TrackSvc is for managing 'tracks' tables in the database.
type TrackSvc struct {
	logic.DbOperations
}

// NewTrackSvc creates a new instance of the TrackSvc struct.
func NewTrackSvc(database logic.DbOperations) TrackSchema {
	return &TrackSvc{database}
}

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (s *TrackSvc) CreateTracksTable() {
	_, err := s.Exec(queries.CreateTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// DropTracksTable drops the tracks table if it exists.
func (s *TrackSvc) DropTracksTable() {
	_, err := s.Exec(queries.DropTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
