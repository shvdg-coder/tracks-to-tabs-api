package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents operations related to references.
type Operations interface {
	database.TabsOps
	mappers.MappingOperations
	GetReferencesCascading(internalID ...uuid.UUID) ([]*models.Reference, error)
}

// ArtistTrackSvc is responsible for managing references.
type Service struct {
	database.TabsOps
	mappers.MappingOperations
	SourcesOps Operations
}

// NewTrackSvc instantiates a new ArtistTrackSvc.
func NewService(data database.TabsOps, mapping mappers.MappingOperations, sources Operations) Operations {
	return &Service{TabsOps: data, MappingOperations: mapping, SourcesOps: sources}
}

// GetReferencesCascading retrieves the Reference's with all their references.
func (s *Service) GetReferencesCascading(internalID ...uuid.UUID) ([]*models.Reference, error) {
	references, err := s.GetReferences(internalID...)
	if err != nil {
		return nil, err
	}

	sourceIDs := s.ExtractSourceIDs(references)
	sources, err := s.SourcesOps.GetSourcesCascading(sourceIDs...)
	if err != nil {
		return nil, err
	}

	sourcesMap := s.SourcesOps.SourcesToMap(sources)
	references = s.MapSourcesToReferences(references, sourcesMap)

	return references, nil
}

// ExtractSourceIDs extracts the source ID from the Reference.
func (s *Service) ExtractSourceIDs(references []*models.Reference) []uint {
	sourceIDMap := make(map[uint]bool)
	for _, reference := range references {
		if reference.Source != nil {
			sourceIDMap[reference.Source.ID] = true
		}
	}
	sourceIDs := make([]uint, 0)
	for key, _ := range sourceIDMap {
		sourceIDs = append(sourceIDs, key)
	}
	return sourceIDs
}
