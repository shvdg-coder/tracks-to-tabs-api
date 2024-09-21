package models

import "encoding/json"

// SourceEntry represents a source in the database.
type SourceEntry struct {
	ID       uint   `yaml:"id"`
	Name     string `yaml:"name"`
	Category string `yaml:"category"`
}

// Fields returns a slice of interfaces containing values of the SourceEntry.
func (s *SourceEntry) Fields() []interface{} {
	return []interface{}{s.ID, s.Name, s.Category}
}

// HasCategory checks if the Source has the provided category.
func (s *SourceEntry) HasCategory(category string) bool {
	return s.Category == category
}

// Source represents a source with its entity references.
type Source struct {
	*SourceEntry
	Endpoints []*Endpoint
}

// MarshalJSON marshals the models.Source.
func (s *Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}
