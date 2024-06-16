package id_references

import "github.com/google/uuid"

// IdReference represents a reference from an internal ID to an external ID.
type IdReference struct {
	InternalSource string
	InternalID     uuid.UUID
	ExternalSource string
	ExternalID     string
}

// NewIdReference instantiates a new IdReference.
func NewIdReference(internalSource string, internalId uuid.UUID, externalSource, externalId string) *IdReference {
	return &IdReference{
		InternalSource: internalSource,
		InternalID:     internalId,
		ExternalSource: externalSource,
		ExternalID:     externalId,
	}
}
