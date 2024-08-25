package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// ReferenceOps represents operations related to references.
type ReferenceOps interface {
	schemas.ReferenceSchema
	data.ReferenceData
	mappers.ReferenceMapper
	GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error)
	ExtractSourceIDs(references []*models.Reference) []uint
}

// ReferenceSvc is responsible for managing references.
type ReferenceSvc struct {
	schemas.ReferenceSchema
	data.ReferenceData
	mappers.ReferenceMapper
	SourceOps
}

// NewReferenceSvc instantiates a new ReferenceSvc.
func NewReferenceSvc(schema schemas.ReferenceSchema, data data.ReferenceData, mapper mappers.ReferenceMapper, sources SourceOps) ReferenceOps {
	return &ReferenceSvc{
		ReferenceSchema: schema,
		ReferenceData:   data,
		ReferenceMapper: mapper,
		SourceOps:       sources,
	}
}

// GetReferences retrieves the models.Reference's with entity references, for the provided internal IDs.
func (r *ReferenceSvc) GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error) {
	referenceEntries, err := r.GetReferenceEntries(internalID...)
	if err != nil {
		return nil, err
	}

	references := r.ReferenceEntriesToReferences(referenceEntries)

	sourceIDs := r.ExtractSourceIDs(references)
	sources, err := r.GetSources(sourceIDs...)
	if err != nil {
		return nil, err
	}

	sourcesMap := r.SourcesToMap(sources)
	references = r.MapSourcesToReferences(references, sourcesMap)

	return references, nil
}

// ExtractSourceIDs extracts the source ID from the models.Reference.
func (r *ReferenceSvc) ExtractSourceIDs(references []*models.Reference) []uint {
	sourceIDMap := make(map[uint]bool)
	for _, reference := range references {
		sourceIDMap[reference.Source.ID] = true
	}

	sourceIDs := make([]uint, 0)
	for key, _ := range sourceIDMap {
		sourceIDs = append(sourceIDs, key)
	}

	return sourceIDs
}
