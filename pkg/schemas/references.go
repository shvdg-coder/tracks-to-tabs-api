package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ReferenceSchema represents schema operations related to 'references' in the database.
type ReferenceSchema interface {
	CreateReferencesTable()
	DropReferencesTable()
}

// ReferenceSvc is for managing 'references' tables in the database.
type ReferenceSvc struct {
	logic.DbOps
}

// NewReferenceSvc creates a new instance of the ReferenceSvc struct.
func NewReferenceSvc(database logic.DbOps) ReferenceSchema {
	return &ReferenceSvc{database}
}

// CreateReferencesTable creates the references table if it doesn't already exist.
func (s *ReferenceSvc) CreateReferencesTable() {
	_, err := s.DB().Exec(queries.CreateReferencesTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropReferencesTable drops the references table if it exists.
func (s *ReferenceSvc) DropReferencesTable() {
	_, err := s.DB().Exec(queries.DropReferencesTable)
	if err != nil {
		log.Fatal(err)
	}
}
