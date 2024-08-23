package mappers

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// SourceMapper represents operations related to source data mapping.
type SourceMapper interface {
	SourcesToMap(sources []*models.Source) map[uint]*models.Source
	MapEndpointsToSources(map[uint]*models.Source, []*models.EndpointEntry) map[uint]*models.Source
}

// SourceSvc is responsible for mapping entities to sources.
type SourceSvc struct {
	SourceMapper
}

// NewSourceSvc creates a new instance of ReferenceSvc.
func NewSourceSvc() SourceMapper {
	return &SourceSvc{}
}

// SourcesToMap transforms a slice of sources into a map where the key is the ID and the value the Source.
func (m *SourceSvc) SourcesToMap(sources []*models.Source) map[uint]*models.Source {
	sourcesMap := make(map[uint]*models.Source)
	for _, source := range sources {
		sourcesMap[source.ID] = source
	}
	return sourcesMap
}

// MapEndpointsToSources maps the endpoints.EndpointEntry's to the Source's.
func (m *SourceSvc) MapEndpointsToSources(sourcesMap map[uint]*models.Source, endpoints []*models.EndpointEntry) map[uint]*models.Source {
	for _, endpoint := range endpoints {
		source := sourcesMap[endpoint.SourceID]
		source.Endpoints = append(source.Endpoints, endpoint)
	}
	return sourcesMap
}
