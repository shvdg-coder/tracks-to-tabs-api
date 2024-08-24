package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// InstrumentSchema represents schema operations related to 'instruments' in the database.
type InstrumentSchema interface {
	CreateInstrumentsTable()
	DropInstrumentsTable()
}

// InstrumentSvc is for managing 'instruments' tables in the database.
type InstrumentSvc struct {
	logic.DbOperations
}

// NewInstrumentSvc creates a new instance of the InstrumentSvc struct.
func NewInstrumentSvc(database logic.DbOperations) InstrumentSchema {
	return &InstrumentSvc{database}
}

// CreateInstrumentsTable creates an instruments table if it doesn't already exist.
func (s *InstrumentSvc) CreateInstrumentsTable() {
	_, err := s.Exec(queries.CreateInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// DropInstrumentsTable drops the instruments table if it exists.
func (s *InstrumentSvc) DropInstrumentsTable() {
	_, err := s.Exec(queries.DropInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
