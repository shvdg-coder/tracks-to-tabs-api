package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// EndpointOps represents operations related to endpoints.
type EndpointOps interface {
	data.EndpointsData
	mappers.EndpointMapper
	GetEndpoints(sourceID ...uint) ([]*models.Endpoint, error)
}

// EndpointSvc is responsible for managing endpoints.
type EndpointSvc struct {
	data.EndpointsData
	mappers.EndpointMapper
}

// NewEndpointSvc instantiates a new EndpointSvc.
func NewEndpointSvc(data data.EndpointsData, mapper mappers.EndpointMapper) EndpointOps {
	return &EndpointSvc{EndpointsData: data, EndpointMapper: mapper}
}

// GetEndpoints retrieves the endpoints, with entity references, for the provided IDs.
func (s *EndpointSvc) GetEndpoints(sourceID ...uint) ([]*models.Endpoint, error) {
	endpointEntries, err := s.GetEndpointEntries(sourceID...)
	if err != nil {
		return nil, err
	}

	endpoints := s.EndpointEntriesToEndpoints(endpointEntries)

	return endpoints, nil
}
