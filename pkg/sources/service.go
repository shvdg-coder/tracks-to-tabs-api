package sources

import end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"

// Operations represents operations related to sources.
type Operations interface {
	DataOperations
	MappingOperations
	GetSourcesCascading(sourceID ...uint) ([]*Source, error)
}

// Service is responsible for managing sources.
type Service struct {
	DataOperations
	MappingOperations
	EndpointsOps end.Operations
}

// NewService instantiates a new Service.
func NewService(data DataOperations, mapping MappingOperations, endpoints end.Operations) Operations {
	return &Service{DataOperations: data, MappingOperations: mapping, EndpointsOps: endpoints}
}

// GetSourcesCascading retrieves all sources with their references.
func (s *Service) GetSourcesCascading(sourceID ...uint) ([]*Source, error) {
	sources, err := s.GetSources(sourceID...)
	if err != nil {
		return nil, err
	}

	endpoints, err := s.EndpointsOps.GetEndpoints(sourceID...)
	if err != nil {
		return nil, err
	}

	sourcesMap := s.SourcesToMap(sources)
	sources = s.MapEndpointsToSources(sourcesMap, endpoints)

	return sources, nil
}
