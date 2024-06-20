package sources

// Source represents a unique resource or provider.
type Source struct {
	ID   uint
	Name string
}

// SourceConfig modifies a Source with configuration options.
type SourceConfig func(*Source)

// WithID sets the ID of a Source.
func WithID(id uint) SourceConfig {
	return func(a *Source) {
		a.ID = id
	}
}

// NewSource instantiates a new Source.
func NewSource(name string, configs ...SourceConfig) *Source {
	source := &Source{Name: name}
	for _, config := range configs {
		config(source)
	}
	return source
}
