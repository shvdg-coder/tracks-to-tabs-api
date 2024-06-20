package resources

import "github.com/google/uuid"

// Resource represents a record in the 'resources' table.
type Resource struct {
	InternalID uuid.UUID
	SourceID   uuid.UUID
	Category   string
	Type       string
	Resource   string
}

// NewResource instantiates a new Resource.
func NewResource(internalId, sourceId uuid.UUID, category, resourceType, resource string) *Resource {
	return &Resource{
		InternalID: internalId,
		SourceID:   sourceId,
		Category:   category,
		Type:       resourceType,
		Resource:   resource,
	}
}
