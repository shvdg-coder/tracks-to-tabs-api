package instruments

// Instrument represents an artist
type Instrument struct {
	ID   uint
	Name string
}

// NewInstrument instantiates a new Instrument.
func NewInstrument(id uint, name string) *Instrument {
	return &Instrument{
		ID:   id,
		Name: name,
	}
}
