package models

import (
	"github.com/google/uuid"
)

// ReferenceEntry represents a reference in the database.
type ReferenceEntry struct {
	InternalID uuid.UUID
	SourceID   uint
	Category   string
	Type       string
	Reference  string
}
