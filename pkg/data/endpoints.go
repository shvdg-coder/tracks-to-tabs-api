package data

import (
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// EndpointsData represents operations related to endpoints in the database.
type EndpointsData interface {
	InsertEndpoints(endpoints ...*models.EndpointEntry)
	InsertEndpoint(endpoint *models.EndpointEntry)
	GetEndpoint(sourceID uint) (*models.EndpointEntry, error)
	GetEndpoints(sourceID ...uint) ([]*models.EndpointEntry, error)
}

// EndpointSvc is for managing endpoints.
type EndpointSvc struct {
	logic.DbOperations
}

// NewEndpointSvc creates a new instance of EndpointSvc.
func NewEndpointSvc(database logic.DbOperations) EndpointsData {
	return &EndpointSvc{DbOperations: database}
}

// InsertEndpoints inserts multiple records into the endpoints table.
func (d *EndpointSvc) InsertEndpoints(endpoints ...*models.EndpointEntry) {
	for _, endpoint := range endpoints {
		d.InsertEndpoint(endpoint)
	}
}

// InsertEndpoint inserts a record into the endpoints table.
func (d *EndpointSvc) InsertEndpoint(endpoint *models.EndpointEntry) {
	_, err := d.Exec(queries.InsertEndpoint, endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.UnformattedURL)
	if err != nil {
		log.Printf(
			"Failed to insert endpoint with Source '%d', Category '%s', Type '%s', and UnformattedURL '%s': %s",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.UnformattedURL, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted endpoint into the 'endpoints' table with Source '%d', Category '%s', Type '%s', and UnformattedURL '%s'",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.UnformattedURL,
		)
	}
}

// GetEndpoint retrieves the endpoint for the provided ID from the database.
func (d *EndpointSvc) GetEndpoint(sourceID uint) (*models.EndpointEntry, error) {
	endpoints, err := d.GetEndpoints(sourceID)
	if err != nil {
		return nil, err
	}
	return endpoints[0], nil
}

// GetEndpoints retrieves the endpoints for the provided IDs from the database.
func (d *EndpointSvc) GetEndpoints(sourceID ...uint) ([]*models.EndpointEntry, error) {
	rows, err := d.Query(queries.GetEndpointsFromIDs, pq.Array(sourceID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var endpoints []*models.EndpointEntry
	for rows.Next() {
		endpoint := &models.EndpointEntry{}
		err := rows.Scan(&endpoint.SourceID, &endpoint.Category, &endpoint.Type, &endpoint.UnformattedURL)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return endpoints, nil
}
