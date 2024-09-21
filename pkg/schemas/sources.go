package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// SourceSchema represents schema operations related to 'sources' in the database.
type SourceSchema interface {
	CreateSourcesTable()
	DropSourcesTable()
}

// SourceSvc is for managing 'sources' tables in the database.
type SourceSvc struct {
	logic.DbOps
}

// NewSourceSvc creates a new instance of the SourceSvc struct.
func NewSourceSvc(database logic.DbOps) SourceSchema {
	return &SourceSvc{database}
}

// CreateSourcesTable creates a sources table if it doesn't already exist.
func (s *SourceSvc) CreateSourcesTable() {
	_, err := s.DB().Exec(queries.CreateSourcesTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropSourcesTable drops the sources table if it exists.
func (s *SourceSvc) DropSourcesTable() {
	_, err := s.DB().Exec(queries.DropSourcesTable)
	if err != nil {
		log.Fatal(err)
	}
}
