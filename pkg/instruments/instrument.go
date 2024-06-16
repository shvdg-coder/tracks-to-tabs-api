package instruments

// Instrument represents an playable instrument.
type Instrument struct {
	ID   uint
	Name string
}

// InstrumentConfig modifies an Instrument with configuration options.
type InstrumentConfig func(*Instrument)

// WithID sets the ID of an Instrument.
func WithID(id uint) InstrumentConfig {
	return func(a *Instrument) {
		a.ID = id
	}
}

// NewInstrument instantiates a new Instrument.
func NewInstrument(name string, configs ...InstrumentConfig) *Instrument {
	instrument := &Instrument{Name: name}
	for _, config := range configs {
		config(instrument)
	}
	return instrument
}
