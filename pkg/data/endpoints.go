package data

import (
	"github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// EndpointsData represents operations related to endpoints in the database.
type EndpointsData interface {
	InsertEndpointEntries(endpoints ...*models.EndpointEntry) error
	GetEndpointEntries(sourceID ...uint) ([]*models.EndpointEntry, error)
	GetEndpointEntry(sourceID uint) (*models.EndpointEntry, error)
}

// EndpointSvc is for managing endpoints.
type EndpointSvc struct {
	logic.DbOps
}

// NewEndpointSvc creates a new instance of EndpointSvc.
func NewEndpointSvc(database logic.DbOps) EndpointsData {
	return &EndpointSvc{DbOps: database}
}

// InsertEndpointEntries inserts multiple records into the endpoints table.
func (d *EndpointSvc) InsertEndpointEntries(endpoints ...*models.EndpointEntry) error {
	data := make([][]interface{}, len(endpoints))

	for i, endpoint := range endpoints {
		data[i] = endpoint.Fields()
	}

	fieldNames := []string{"source_id", "category", "type", "url"}
	return d.BulkInsert("endpoints", fieldNames, data)
}

// GetEndpointEntry retrieves the endpoint for the provided ID from the database.
func (d *EndpointSvc) GetEndpointEntry(sourceID uint) (*models.EndpointEntry, error) {
	endpoints, err := d.GetEndpointEntries(sourceID)
	if err != nil {
		return nil, err
	}
	return endpoints[0], nil
}

// GetEndpointEntries retrieves the endpoints for the provided IDs from the database.
func (d *EndpointSvc) GetEndpointEntries(sourceID ...uint) ([]*models.EndpointEntry, error) {
	rows, err := d.DB().Query(queries.GetEndpointsFromIDs, pq.Array(sourceID))
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
