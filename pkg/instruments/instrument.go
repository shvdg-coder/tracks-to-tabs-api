package instruments

// Instrument represents an playable instrument.
type Instrument struct {
	ID   uint   `yaml:"id"`
	Name string `yaml:"name"`
}

type Option func(*Instrument)

// WithID sets the ID of an instrument
func WithID(id uint) Option {
	return func(i *Instrument) {
		i.ID = id
	}
}

// NewInstrument instantiates a new Instrument.
func NewInstrument(name string, options ...Option) *Instrument {
	instrument := &Instrument{Name: name}
	for _, option := range options {
		option(instrument)
	}
	return instrument
}
