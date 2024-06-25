package sources

// Source represents a unique provider.
type Source struct {
	ID         uint     `yaml:"id"`
	Name       string   `yaml:"name"`
	Categories []string `yaml:"categories"`
}

// SourceOptional sets an optional value of the Source object.
type SourceOptional func(source *Source)

// WithCategories sets the categories property of the Source object.
func WithCategories(categories []string) SourceOptional {
	return func(source *Source) {
		source.Categories = categories
	}
}

// NewSource instantiates a new Source.
func NewSource(id uint, name string, options ...SourceOptional) *Source {
	source := &Source{ID: id, Name: name}
	for _, option := range options {
		option(source)
	}
	return source
}

// HasCategory checks if the Source has the provided category.
func (s *Source) HasCategory(category string) bool {
	for _, cat := range s.Categories {
		if cat == category {
			return true
		}
	}
	return false
}

// HasCategories checks if the Source contains all the provided categories.
func (s *Source) HasCategories(categories ...string) bool {
	for _, requiredCat := range categories {
		if !s.HasCategory(requiredCat) {
			return false
		}
	}
	return true
}
