package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// SourceOps represents operations related to sources.
type SourceOps interface {
	data.SourceData
	mappers.SourceMapper
	GetSources(sourceID ...uint) ([]*models.Source, error)
}

// SourceSvc is responsible for managing sources.
type SourceSvc struct {
	data.SourceData
	mappers.SourceMapper
	EndpointOps
}

// NewSourceSvc instantiates a new SourceSvc.
func NewSourceSvc(data data.SourceData, mapper mappers.SourceMapper, endpoints EndpointOps) SourceOps {
	return &SourceSvc{SourceData: data, SourceMapper: mapper, EndpointOps: endpoints}
}

// GetSources retrieves sources with their entity references.
func (s *SourceSvc) GetSources(sourceID ...uint) ([]*models.Source, error) {
	sourceEntries, err := s.GetSourceEntries(sourceID...)
	if err != nil {
		return nil, err
	}

	endpoints, err := s.GetEndpoints(sourceID...)
	if err != nil {
		return nil, err
	}

	sources := s.SourceEntriesToSources(sourceEntries)
	sourcesMap := s.SourcesToMap(sources)
	sourcesMap = s.MapEndpointsToSources(sourcesMap, endpoints)
	sources = s.MapToSources(sourcesMap)

	return sources, nil
}
