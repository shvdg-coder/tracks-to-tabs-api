package models

// Source represents a source with its entity references.
type Source struct {
	*SourceEntry
	Endpoints []*EndpointEntry
}

// HasCategory checks if the Source has the provided category.
func (s *Source) HasCategory(category string) bool {
	return s.Category == category
}
