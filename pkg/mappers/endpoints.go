package mappers

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"

// EndpointMapper represents operations related to endpoint data mapping.
type EndpointMapper interface {
	EndpointEntriesToEndpoints(endpointEntries []*models.EndpointEntry) []*models.Endpoint
}

// EndpointSvc is responsible for mapping entities to endpoints.
type EndpointSvc struct {
	EndpointMapper
}

// NewEndpointSvc creates a new instance of EndpointSvc.
func NewEndpointSvc() EndpointMapper {
	return &EndpointSvc{}
}

// EndpointEntriesToEndpoints transforms a slice of EndpointEntry's into a slice of Endpoint's.
func (e *EndpointSvc) EndpointEntriesToEndpoints(endpointEntries []*models.EndpointEntry) []*models.Endpoint {
	endpoints := make([]*models.Endpoint, len(endpointEntries))
	for i, endpointEntry := range endpointEntries {
		endpoints[i] = &models.Endpoint{EndpointEntry: endpointEntry, Source: &models.Source{SourceEntry: &models.SourceEntry{ID: endpointEntry.SourceID}}}
	}
	return endpoints
}
