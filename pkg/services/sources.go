package services

import (
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/schemas"
)

// SourceOps represents operations related to sources.
type SourceOps interface {
	schemas.SourceSchema
	data.SourceData
	mappers.SourceMapper
	GetSources(sourceID ...uint) ([]*models.Source, error)
	GetSourcesCascading(sourceID ...uint) ([]*models.Source, error)
}

// SourceSvc is responsible for managing sources.
type SourceSvc struct {
	schemas.SourceSchema
	data.SourceData
	mappers.SourceMapper
	EndpointOps
}

// NewSourceSvc instantiates a new SourceSvc.
func NewSourceSvc(schema schemas.SourceSchema, data data.SourceData, mapper mappers.SourceMapper, endpoints EndpointOps) SourceOps {
	return &SourceSvc{
		SourceSchema: schema,
		SourceData:   data,
		SourceMapper: mapper,
		EndpointOps:  endpoints,
	}
}

// GetSources retrieves sources without their entity references.
func (s *SourceSvc) GetSources(sourceID ...uint) ([]*models.Source, error) {
	sourceEntries, err := s.GetSourceEntries(sourceID...)
	if err != nil {
		return nil, err
	}

	sources := s.SourceEntriesToSources(sourceEntries)

	return sources, nil
}

// GetSourcesCascading retrieves sources with their entity references.
func (s *SourceSvc) GetSourcesCascading(sourceID ...uint) ([]*models.Source, error) {
	sources, err := s.GetSources(sourceID...)
	if err != nil {
		return nil, err
	}

	err = s.LoadEndpoints(sources...)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

// LoadEndpoints loads the models.Endpoint's for given models.Source's.
func (s *SourceSvc) LoadEndpoints(sources ...*models.Source) error {
	sourceIDs := s.ExtractIDsFromSources(sources)
	endpoints, err := s.GetEndpointsCascading(sourceIDs...)
	if err != nil {
		return err
	}

	sourcesMap := s.SourcesToMap(sources)
	s.MapSourcesToEndpoints(endpoints, sourcesMap)
	s.MapEndpointsToSources(sourcesMap, endpoints)

	return nil
}

// ExtractIDsFromSources retrieves the ID's from the models.Source's.
func (s *SourceSvc) ExtractIDsFromSources(sources []*models.Source) []uint {
	sourceIDs := make([]uint, len(sources))
	for i, source := range sources {
		sourceIDs[i] = source.ID
	}
	return sourceIDs
}
