package services

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
)

// EndpointOps represents operations related to endpoints.
type EndpointOps interface {
	schemas.EndpointSchema
	data.EndpointsData
	mappers.EndpointMapper
	GetEndpoints(sourceID ...uint) ([]*models.Endpoint, error)
	GetEndpointsCascading(sourceID ...uint) ([]*models.Endpoint, error)
	ExtractIDsFromEndpoints(endpoints []*models.Endpoint) []uint
}

// EndpointSvc is responsible for managing endpoints.
type EndpointSvc struct {
	schemas.EndpointSchema
	data.EndpointsData
	mappers.EndpointMapper
}

// NewEndpointSvc instantiates a new EndpointSvc.
func NewEndpointSvc(schema schemas.EndpointSchema, data data.EndpointsData, mapper mappers.EndpointMapper) EndpointOps {
	return &EndpointSvc{
		EndpointSchema: schema,
		EndpointsData:  data,
		EndpointMapper: mapper,
	}
}

// GetEndpoints retrieves the endpoints, with entity references, for the provided IDs.
func (e *EndpointSvc) GetEndpoints(sourceID ...uint) ([]*models.Endpoint, error) {
	endpointEntries, err := e.GetEndpointEntries(sourceID...)
	if err != nil {
		return nil, err
	}

	endpoints := e.EndpointEntriesToEndpoints(endpointEntries)

	return endpoints, nil
}

// GetEndpointsCascading retrieves the endpoints, with entity references, for the provided IDs.
func (e *EndpointSvc) GetEndpointsCascading(sourceID ...uint) ([]*models.Endpoint, error) {
	endpoints, err := e.GetEndpoints(sourceID...)
	if err != nil {
		return nil, err
	}
	return endpoints, nil
}

// ExtractIDsFromEndpoints retrieves the source IDs from the models.Endpoint's.
func (e *EndpointSvc) ExtractIDsFromEndpoints(endpoints []*models.Endpoint) []uint {
	sourceIDs := make([]uint, len(endpoints))
	for i, endpoint := range endpoints {
		sourceIDs[i] = endpoint.Source.ID
	}
	return sourceIDs
}
