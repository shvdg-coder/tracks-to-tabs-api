package references

import "github.com/google/uuid"

// Reference represents a record in the 'references' table.
type Reference struct {
	InternalID uuid.UUID
	SourceID   uint
	Category   string
	Type       string
	Reference  string
}

// NewReference instantiates a new Reference.
func NewReference(internalId uuid.UUID, sourceId uint, category, referenceType, reference string) *Reference {
	return &Reference{
		InternalID: internalId,
		SourceID:   sourceId,
		Category:   category,
		Type:       referenceType,
		Reference:  reference,
	}
}
