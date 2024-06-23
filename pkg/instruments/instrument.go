package instruments

// Instrument represents an playable instrument.
type Instrument struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

// NewInstrument instantiates a new Instrument.
func NewInstrument(id uint, name string) *Instrument {
	return &Instrument{ID: id, Name: name}
}
