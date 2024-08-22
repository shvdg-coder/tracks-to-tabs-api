package models

// SourceEntry represents a source in the database.
type SourceEntry struct {
	ID       uint   `yaml:"id"`
	Name     string `yaml:"name"`
	Category string `yaml:"category"`
}
