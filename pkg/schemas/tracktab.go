package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TrackTabSchema represents schema operations related to 'track_tab' links in the database.
type TrackTabSchema interface {
	CreateTrackTabTable()
	DropTrackTabTable()
}

// TrackTabSvc is for managing 'track_tab' links in the database.
type TrackTabSvc struct {
	logic.DbOperations
}

// NewTrackTabSvc creates a new instance of the TrackTabSvc struct.
func NewTrackTabSvc(database logic.DbOperations) TrackTabSchema {
	return &TrackTabSvc{database}
}

// CreateTrackTabTable creates the track_tab table if it doesn't already exist.
func (s *TrackTabSvc) CreateTrackTabTable() {
	_, err := s.Exec(queries.CreateTrackTabTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropTrackTabTable drops the track_tab table if it exists.
func (s *TrackTabSvc) DropTrackTabTable() {
	_, err := s.Exec(queries.DropTrackTabTable)
	if err != nil {
		log.Fatal(err)
	}
}
