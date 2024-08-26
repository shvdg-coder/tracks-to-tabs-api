package mappers

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"

// EndpointMapper represents operations related to endpoint data mapping.
type EndpointMapper interface {
	MapSourcesToEndpoints(endpoints []*models.Endpoint, sources map[uint]*models.Source) []*models.Endpoint
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

// MapSourcesToEndpoints maps the models.Source's to the models.Endpoint's, by updating the provided models.Endpoint slice and returning it.
func (e *EndpointSvc) MapSourcesToEndpoints(endpoints []*models.Endpoint, sources map[uint]*models.Source) []*models.Endpoint {
	for _, endpoint := range endpoints {
		source := sources[endpoint.SourceID]
		endpoint.Source = source
	}
	return endpoints
}

// EndpointEntriesToEndpoints transforms a slice of EndpointEntry's into a slice of Endpoint's.
func (e *EndpointSvc) EndpointEntriesToEndpoints(endpointEntries []*models.EndpointEntry) []*models.Endpoint {
	endpoints := make([]*models.Endpoint, len(endpointEntries))
	for i, endpointEntry := range endpointEntries {
		endpoints[i] = &models.Endpoint{EndpointEntry: endpointEntry, Source: &models.Source{SourceEntry: &models.SourceEntry{ID: endpointEntry.SourceID}}}
	}
	return endpoints
}
