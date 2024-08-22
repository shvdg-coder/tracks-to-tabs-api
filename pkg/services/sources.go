package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents operations related to sources.
type Operations interface {
	database.TabsOps
	mappers.MappingOperations
	GetSourcesCascading(sourceID ...uint) ([]*models.Source, error)
}

// ArtistTrackSvc is responsible for managing sources.
type Service struct {
	database.TabsOps
	mappers.MappingOperations
	EndpointsOps Operations
}

// NewTrackSvc instantiates a new ArtistTrackSvc.
func NewService(data database.TabsOps, mapping mappers.MappingOperations, endpoints Operations) Operations {
	return &Service{TabsOps: data, MappingOperations: mapping, EndpointsOps: endpoints}
}

// GetSourcesCascading retrieves all sources with their references.
func (s *Service) GetSourcesCascading(sourceID ...uint) ([]*models.Source, error) {
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
