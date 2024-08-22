package mappers

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// MappingOperations represents operations related to reference data mapping.
type MappingOperations interface {
	MapSourcesToReferences(references []*models.Reference, sourcesMap map[uint]*models.Source) []*models.Reference
}

// MappingService is responsible for mapping entities to references.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// MapSourcesToReferences maps the sources.Source's to the Reference's.
func (m *MappingService) MapSourcesToReferences(references []*models.Reference, sourcesMap map[uint]*models.Source) []*models.Reference {
	for _, reference := range references {
		if reference.Source == nil {
			continue
		}
		source, ok := sourcesMap[reference.Source.ID]
		if !ok {
			continue
		}
		reference.Source = source
	}
	var referencesResult []*references.Reference
	for _, reference := range references {
		referencesResult = append(referencesResult, reference)
	}
	return referencesResult
}
