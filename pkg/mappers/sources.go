package mappers

import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
)

// SourceMapper represents operations related to source data mapping.
type SourceMapper interface {
	SourceEntriesToSources(sources []*models.SourceEntry) []*models.Source
	SourcesToMap(sources []*models.Source) map[uint]*models.Source
	MapToSources(sourcesMap map[uint]*models.Source) []*models.Source
	MapEndpointsToSources(map[uint]*models.Source, []*models.Endpoint) map[uint]*models.Source
}

// SourceSvc is responsible for mapping entities to sources.
type SourceSvc struct {
	SourceMapper
}

// NewSourceSvc creates a new instance of ReferenceSvc.
func NewSourceSvc() SourceMapper {
	return &SourceSvc{}
}

// SourceEntriesToSources transforms a slice of models.SourceEntry's into a slice of []*models.Source.
func (s *SourceSvc) SourceEntriesToSources(sourceEntries []*models.SourceEntry) []*models.Source {
	sources := make([]*models.Source, len(sourceEntries))
	for i, sourceEntry := range sourceEntries {
		sources[i] = &models.Source{SourceEntry: sourceEntry}
	}
	return sources
}

// SourcesToMap transforms a slice of models.Source's into a map where the key is the ID and the value the models.Source.
func (s *SourceSvc) SourcesToMap(sources []*models.Source) map[uint]*models.Source {
	sourcesMap := make(map[uint]*models.Source, len(sources))
	for _, source := range sources {
		sourcesMap[source.ID] = source
	}
	return sourcesMap
}

// MapToSources transforms a map of models.Source's into a slice.
func (s *SourceSvc) MapToSources(sourcesMap map[uint]*models.Source) []*models.Source {
	sources := make([]*models.Source, 0)
	for _, source := range sourcesMap {
		sources = append(sources, source)
	}
	return sources
}

// MapEndpointsToSources adds the models.Endpoint's to the models.Source's, by updating the provided models.Source's map and returning it.
func (s *SourceSvc) MapEndpointsToSources(sourcesMap map[uint]*models.Source, endpoints []*models.Endpoint) map[uint]*models.Source {
	for _, endpoint := range endpoints {
		source := sourcesMap[endpoint.SourceID]
		source.Endpoints = append(source.Endpoints, endpoint)
	}
	return sourcesMap
}
