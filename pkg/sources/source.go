package sources

// Source represents a unique resource or provider.
type Source struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// NewSource instantiates a new Source.
func NewSource(id uint, name string) *Source {
	return &Source{ID: id, Name: name}
}
