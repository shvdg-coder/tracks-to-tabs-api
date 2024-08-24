package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// EndpointSchema represents schema operations related to 'endpoints' in the database.
type EndpointSchema interface {
	CreateEndpointsTable()
	DropEndpointsTable()
}

// EndpointSvc is for managing 'endpoints' tables in the database.
type EndpointSvc struct {
	logic.DbOperations
}

// NewEndpointSvc creates a new instance of the EndpointSvc struct.
func NewEndpointSvc(database logic.DbOperations) EndpointSchema {
	return &EndpointSvc{database}
}

// CreateEndpointsTable creates the endpoints table if it doesn't already exist.
func (s *EndpointSvc) CreateEndpointsTable() {
	_, err := s.Exec(queries.CreateEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

// DropEndpointsTable drops the endpoints table if it exists.
func (s *EndpointSvc) DropEndpointsTable() {
	_, err := s.Exec(queries.DropEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
