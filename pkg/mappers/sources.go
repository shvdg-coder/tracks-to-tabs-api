package mappers

import (
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// MappingOperations represents operations related to source data mapping.
type MappingOperations interface {
	SourcesToMap(sources []*end.Source) map[uint]*end.Source
	MapEndpointsToSources(map[uint]*end.Source, []*end.EndpointEntry) []*end.Source
}

// MappingService is responsible for mapping entities to sources.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// SourcesToMap transforms a slice of sources into a map where the key is the ID and the value the Source.
func (m *MappingService) SourcesToMap(sources []*end.Source) map[uint]*end.Source {
	sourcesMap := make(map[uint]*sources.Source)
	for _, source := range sources {
		sourcesMap[source.ID] = source
	}
	return sourcesMap
}

// MapEndpointsToSources maps the endpoints.EndpointEntry's to the Source's.
func (m *MappingService) MapEndpointsToSources(sourcesMap map[uint]*end.Source, endpoints []*end.EndpointEntry) []*end.Source {
	for _, endpoint := range endpoints {
		source, ok := sourcesMap[endpoint.SourceID]
		if ok {
			source.Endpoints = append(source.Endpoints, endpoint)
		}
	}
	var sources []*end.Source
	for _, source := range sourcesMap {
		sources = append(sources, source)
	}
	return sources
}
