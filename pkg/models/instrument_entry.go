package models

// InstrumentEntry represents an instrument in the database.
type InstrumentEntry struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}
