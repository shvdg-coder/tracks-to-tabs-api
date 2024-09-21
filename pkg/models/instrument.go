package models

import "encoding/json"

// InstrumentEntry represents an instrument in the database.
type InstrumentEntry struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// Fields returns a slice of interfaces containing values of the InstrumentEntry.
func (i *InstrumentEntry) Fields() []interface{} {
	return []interface{}{i.ID, i.Name}
}

// Instrument represents an instrument with entity references.
type Instrument struct {
	*InstrumentEntry
}

// MarshalJSON marshals the models.Instrument.
func (i *Instrument) MarshalJSON() ([]byte, error) {
	return json.Marshal(*i)
}
