package models

import "encoding/json"

// SourceEntry represents a source in the database.
type SourceEntry struct {
	ID       uint   `yaml:"id"`
	Name     string `yaml:"name"`
	Category string `yaml:"category"`
}

// Source represents a source with its entity references.
type Source struct {
	*SourceEntry
	Endpoints []*Endpoint
}

// HasCategory checks if the Source has the provided category.
func (s *SourceEntry) HasCategory(category string) bool {
	return s.Category == category
}

// MarshalJSON marshals the models.Source.
func (s *Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}
