package sources

import end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"

// Source represents a unique provider.
type Source struct {
	ID        uint   `yaml:"id"`
	Name      string `yaml:"name"`
	Category  string `yaml:"category"`
	Endpoints []*end.Endpoint
}

// SourceOptional sets an optional value of the Source object.
type SourceOptional func(source *Source)

// WithCategory sets the category property of the Source object.
func WithCategory(category string) SourceOptional {
	return func(source *Source) {
		source.Category = category
	}
}

// WithEndpoint sets the endpoints property of the Source object.
func WithEndpoint(endpoints *end.Endpoint) SourceOptional {
	return func(source *Source) {
		source.Endpoints = append(source.Endpoints, endpoints)
	}
}

// NewSource instantiates a new Source.
func NewSource(id uint, name string, options ...SourceOptional) *Source {
	source := &Source{ID: id, Name: name}
	source.Endpoints = make([]*end.Endpoint, 0)
	for _, option := range options {
		option(source)
	}
	return source
}

// HasCategory checks if the Source has the provided category.
func (s *Source) HasCategory(category string) bool {
	return s.Category == category
}
