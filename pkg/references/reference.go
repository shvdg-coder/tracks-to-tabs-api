package references

import (
	"github.com/google/uuid"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
)

// Reference represents a record in the 'references' table.
type Reference struct {
	InternalID uuid.UUID
	Source     *src.Source
	Category   string
	Type       string
	Reference  string
}

// NewReference instantiates a new Reference.
func NewReference(internalId uuid.UUID, source *src.Source, category, referenceType, reference string) *Reference {
	return &Reference{
		InternalID: internalId,
		Source:     source,
		Category:   category,
		Type:       referenceType,
		Reference:  reference,
	}
}
