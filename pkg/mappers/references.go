package mappers

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ReferenceMapper represents operations related to reference data mapping.
type ReferenceMapper interface {
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

// MapSourcesToReferences maps the sources.Source's to the Reference's.
func (r *ReferenceSvc) MapSourcesToReferences(references []*models.Reference, sourcesMap map[uint]*models.Source) []*models.Reference {
	for _, reference := range references {
		source := sourcesMap[reference.Source.ID]
		reference.Source = source
	}
	return references
}
