package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// ReferenceEntry represents a reference in the database.
type ReferenceEntry struct {
	InternalID uuid.UUID `db:"internal_id"`
	SourceID   uint      `db:"source_id"`
	Category   string    `db:"category"`
	Type       string    `db:"type"`
	Reference  string    `db:"reference"`
}

// Fields returns a slice of interfaces containing values of the ReferenceEntry.
func (r *ReferenceEntry) Fields() []interface{} {
	return []interface{}{r.InternalID, r.SourceID, r.Category, r.Type, r.Reference}
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
