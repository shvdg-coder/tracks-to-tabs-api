package models

import (
	"encoding/json"
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

// Reference represents a record in the 'references' table.
type Reference struct {
	*ReferenceEntry
	Source *Source
}

// MarshalJSON marshals the models.Reference while preventing circling.
func (r *Reference) MarshalJSON() ([]byte, error) {
	reference := *r
	reference.Source = &Source{
		SourceEntry: r.Source.SourceEntry,
		Endpoints:   r.Source.Endpoints,
	}
	return json.Marshal(reference)
}
