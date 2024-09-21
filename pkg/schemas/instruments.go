package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// InstrumentSchema represents schema operations related to 'instruments' in the database.
type InstrumentSchema interface {
	CreateInstrumentsTable()
	DropInstrumentsTable()
}

// InstrumentSvc is for managing 'instruments' tables in the database.
type InstrumentSvc struct {
	logic.DbOps
}

// NewInstrumentSvc creates a new instance of the InstrumentSvc struct.
func NewInstrumentSvc(database logic.DbOps) InstrumentSchema {
	return &InstrumentSvc{database}
}

// CreateInstrumentsTable creates an instruments table if it doesn't already exist.
func (s *InstrumentSvc) CreateInstrumentsTable() {
	_, err := s.DB().Exec(queries.CreateInstrumentsTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropInstrumentsTable drops the instruments table if it exists.
func (s *InstrumentSvc) DropInstrumentsTable() {
	_, err := s.DB().Exec(queries.DropInstrumentsTable)
	if err != nil {
		log.Fatal(err)
	}
}
