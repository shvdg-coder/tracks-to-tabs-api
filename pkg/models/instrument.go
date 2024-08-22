package models

// InstrumentEntry represents an instrument in the database.
type InstrumentEntry struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// Instrument represents an instrument with entity references.
type Instrument struct {
	*InstrumentEntry
}
