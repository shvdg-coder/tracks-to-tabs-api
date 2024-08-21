package instruments

// Operations represents the operations related to instruments.
type Operations interface {
	DataOperations
	InstrumentsToMap(instruments []*Instrument) map[uint]*Instrument
}

// Service is responsible for managing instruments.
type Service struct {
	DataOperations
}

// NewService creates a new instance of Service.
func NewService(data DataOperations) Operations {
	return &Service{DataOperations: data}
}

// InstrumentsToMap transforms a slice of Instrument's into map, where the key is the ID and the value the Instrument.
func (s *Service) InstrumentsToMap(instruments []*Instrument) map[uint]*Instrument {
	instrumentMap := make(map[uint]*Instrument)
	for _, instrument := range instruments {
		instrumentMap[instrument.ID] = instrument
	}
	return instrumentMap
}
