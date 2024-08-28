package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// DifficultySchema represents schema operations related to difficulties in the database.
type DifficultySchema interface {
	CreateDifficultiesTable()
	DropDifficultiesTable()
}

// DifficultySvc is for managing difficulties.
type DifficultySvc struct {
	logic.DbOperations
}

// NewDifficultySvc creates a new instance of the DifficultySvc struct.
func NewDifficultySvc(database logic.DbOperations) DifficultySchema {
	return &DifficultySvc{database}
}

// CreateDifficultiesTable creates a difficulties table if it doesn't already exist.
func (s *DifficultySvc) CreateDifficultiesTable() {
	_, err := s.Exec(queries.CreateDifficultiesTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropDifficultiesTable drops the difficulties table if it exists.
func (s *DifficultySvc) DropDifficultiesTable() {
	_, err := s.Exec(queries.DropDifficultiesTable)
	if err != nil {
		log.Fatal(err)
	}
}
