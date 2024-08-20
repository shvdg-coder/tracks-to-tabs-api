package endpoints

import (
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to endpoints in the database.
type DataOperations interface {
	InsertEndpoints(endpoints ...*Endpoint)
	InsertEndpoint(endpoint *Endpoint)
	GetEndpoint(sourceID uint) (*Endpoint, error)
	GetEndpoints(sourceID ...uint) ([]*Endpoint, error)
}

// DataService is for managing endpoints.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of DataService.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertEndpoints inserts multiple records into the endpoints table.
func (d *DataService) InsertEndpoints(endpoints ...*Endpoint) {
	for _, endpoint := range endpoints {
		d.InsertEndpoint(endpoint)
	}
}

// InsertEndpoint inserts a record into the endpoints table.
func (d *DataService) InsertEndpoint(endpoint *Endpoint) {
	_, err := d.Exec(insertEndpointQuery, endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL)
	if err != nil {
		log.Printf(
			"Failed to insert endpoint with Source '%d', Category '%s', Type '%s', and URL '%s': %s",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted endpoint into the 'endpoints' table with Source '%d', Category '%s', Type '%s', and URL '%s'",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL,
		)
	}
}

// GetEndpoint retrieves the endpoint for the provided ID from the database.
func (d *DataService) GetEndpoint(sourceID uint) (*Endpoint, error) {
	endpoints, err := d.GetEndpoints(sourceID)
	if err != nil {
		return nil, err
	}
	return endpoints[0], nil
}

// GetEndpoints retrieves the endpoints for the provided IDs from the database.
func (d *DataService) GetEndpoints(sourceID ...uint) ([]*Endpoint, error) {
	rows, err := d.Query(getEndpointsFromIDs, pq.Array(sourceID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var endpoints []*Endpoint
	for rows.Next() {
		endpoint := &Endpoint{}
		err := rows.Scan(&endpoint.SourceID, &endpoint.Category, &endpoint.Type, &endpoint.URL)
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
