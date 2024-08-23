package mappers

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"

// InstrumentMapper represents operations related to instrument data mapping.
type InstrumentMapper interface {
	InstrumentsToMap(instruments []*models.InstrumentEntry) map[uint]*models.InstrumentEntry
}

// InstrumentSvc is responsible for mapping entities to instruments.
type InstrumentSvc struct {
	InstrumentMapper
}

// NewInstrumentSvc creates a new instance of ReferenceSvc.
func NewInstrumentSvc() InstrumentMapper {
	return &InstrumentSvc{}
}

// InstrumentsToMap transforms a slice of InstrumentEntry's into map, where the key is the ID and the value the InstrumentEntry.
func (i *InstrumentSvc) InstrumentsToMap(instruments []*models.InstrumentEntry) map[uint]*models.InstrumentEntry {
	instrumentMap := make(map[uint]*models.InstrumentEntry)
	for _, instrument := range instruments {
		instrumentMap[instrument.ID] = instrument
	}
	return instrumentMap
}
