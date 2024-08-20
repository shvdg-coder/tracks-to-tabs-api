package sources

import end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"

// Operations represents operations related to sources.
type Operations interface {
	DataOperations
	GetSourcesCascading(sourceID ...uint) ([]*Source, error)
}

// Service is responsible for managing sources.
type Service struct {
	DataOperations
	EndpointsOps end.Operations
}

// NewService instantiates a new Service.
func NewService(data DataOperations, endpoints end.Operations) Operations {
	return &Service{DataOperations: data, EndpointsOps: endpoints}
}

// GetSourcesCascading todo:
func (s *Service) GetSourcesCascading(sourceID ...uint) ([]*Source, error) {
	sources, err := s.GetSources(sourceID...)
	if err != nil {
		return nil, err
	}

	_, err = s.EndpointsOps.GetEndpoints(sourceID...)
	if err != nil {
		return nil, err
	}
	//TODO: use endpoints to complete sources objects

	return sources, nil
}
