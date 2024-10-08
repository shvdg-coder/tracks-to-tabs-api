package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TrackTabSchema represents schema operations related to 'track_tab' links in the database.
type TrackTabSchema interface {
	CreateTrackTabTable()
	DropTrackTabTable()
}

// TrackTabSvc is for managing 'track_tab' links in the database.
type TrackTabSvc struct {
	logic.DbOps
}

// NewTrackTabSvc creates a new instance of the TrackTabSvc struct.
func NewTrackTabSvc(database logic.DbOps) TrackTabSchema {
	return &TrackTabSvc{database}
}

// CreateTrackTabTable creates the track_tab table if it doesn't already exist.
func (s *TrackTabSvc) CreateTrackTabTable() {
	_, err := s.DB().Exec(queries.CreateTrackTabTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropTrackTabTable drops the track_tab table if it exists.
func (s *TrackTabSvc) DropTrackTabTable() {
	_, err := s.DB().Exec(queries.DropTrackTabTable)
	if err != nil {
		log.Fatal(err)
	}
}
