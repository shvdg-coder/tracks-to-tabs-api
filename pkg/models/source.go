package models

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
func (s *Source) HasCategory(category string) bool {
	return s.Category == category
}
