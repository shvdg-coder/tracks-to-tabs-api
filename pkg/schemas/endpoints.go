package schemas

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// EndpointSchema represents schema operations related to 'endpoints' in the database.
type EndpointSchema interface {
	CreateEndpointsTable()
	DropEndpointsTable()
}

// EndpointSvc is for managing 'endpoints' tables in the database.
type EndpointSvc struct {
	logic.DbOps
}

// NewEndpointSvc creates a new instance of the EndpointSvc struct.
func NewEndpointSvc(database logic.DbOps) EndpointSchema {
	return &EndpointSvc{database}
}

// CreateEndpointsTable creates the endpoints table if it doesn't already exist.
func (s *EndpointSvc) CreateEndpointsTable() {
	_, err := s.DB().Exec(queries.CreateEndpointsTable)
	if err != nil {
		log.Fatal(err)
	}
}

// DropEndpointsTable drops the endpoints table if it exists.
func (s *EndpointSvc) DropEndpointsTable() {
	_, err := s.DB().Exec(queries.DropEndpointsTable)
	if err != nil {
		log.Fatal(err)
	}
}
