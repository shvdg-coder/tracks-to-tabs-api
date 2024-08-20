package sources

import end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"

// MappingOperations represents operations related to source data mapping.
type MappingOperations interface {
	SourcesToMap(sources []*Source) map[uint]*Source
	MapEndpointsToSources(map[uint]*Source, []*end.Endpoint) []*Source
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
func (m *MappingService) SourcesToMap(sources []*Source) map[uint]*Source {
	sourcesMap := make(map[uint]*Source)
	for _, source := range sources {
		sourcesMap[source.ID] = source
	}
	return sourcesMap
}

// MapEndpointsToSources maps the endpoints.Endpoint's to the Source's.
func (m *MappingService) MapEndpointsToSources(sourcesMap map[uint]*Source, endpoints []*end.Endpoint) []*Source {
	for _, endpoint := range endpoints {
		source, ok := sourcesMap[endpoint.SourceID]
		if ok {
			source.Endpoints = append(source.Endpoints, endpoint)
		}
	}
	var sources []*Source
	for _, source := range sourcesMap {
		sources = append(sources, source)
	}
	return sources
}
