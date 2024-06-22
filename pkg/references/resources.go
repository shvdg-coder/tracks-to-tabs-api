package references

import "github.com/google/uuid"

// Reference represents a record in the 'references' table.
type Reference struct {
	InternalID uuid.UUID
	SourceID   uuid.UUID
	Category   string
	Type       string
	Reference  string
}

// NewResource instantiates a new Reference.
func NewResource(internalId, sourceId uuid.UUID, category, referenceType, reference string) *Reference {
	return &Reference{
		InternalID: internalId,
		SourceID:   sourceId,
		Category:   category,
		Type:       referenceType,
		Reference:  reference,
	}
}
