package mappers

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
)

// ReferenceMapper represents operations related to reference data mapping.
type ReferenceMapper interface {
	ReferenceEntriesToReferences(references []*models.ReferenceEntry) []*models.Reference
	MapSourcesToReferences(references []*models.Reference, sourcesMap map[uint]*models.Source) []*models.Reference
}

// ReferenceSvc is responsible for mapping entities to references.
type ReferenceSvc struct {
	ReferenceMapper
}

// NewReferenceSvc creates a new instance of ReferenceSvc.
func NewReferenceSvc() ReferenceMapper {
	return &ReferenceSvc{}
}

// ReferenceEntriesToReferences transforms a slice of models.ReferenceEntry's into a slice of models.Reference's.
func (r *ReferenceSvc) ReferenceEntriesToReferences(references []*models.ReferenceEntry) []*models.Reference {
	result := make([]*models.Reference, len(references))
	for i, referenceEntry := range references {
		result[i] = &models.Reference{ReferenceEntry: referenceEntry, Source: &models.Source{SourceEntry: &models.SourceEntry{ID: referenceEntry.SourceID}}}
	}
	return result
}

// MapSourcesToReferences maps the models.Source's to the models.Reference's, by updating the provided models.Reference slice and returning it.
func (r *ReferenceSvc) MapSourcesToReferences(references []*models.Reference, sourcesMap map[uint]*models.Source) []*models.Reference {
	for _, reference := range references {
		source := sourcesMap[reference.Source.ID]
		reference.Source = source
	}
	return references
}
