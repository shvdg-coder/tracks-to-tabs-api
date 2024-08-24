package mappers

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"

// InstrumentMapper represents operations related to instrument data mapping.
type InstrumentMapper interface {
	InstrumentEntriesToInstruments(instrumentEntries []*models.InstrumentEntry) []*models.Instrument
	InstrumentsToMap(instruments []*models.Instrument) map[uint]*models.Instrument
}

// InstrumentSvc is responsible for mapping entities to instruments.
type InstrumentSvc struct {
	InstrumentMapper
}

// NewInstrumentSvc creates a new instance of ReferenceSvc.
func NewInstrumentSvc() InstrumentMapper {
	return &InstrumentSvc{}
}

// InstrumentEntriesToInstruments transforms a slice of models.InstrumentEntry's into a slice of models.Instrument's.
func (i *InstrumentSvc) InstrumentEntriesToInstruments(instrumentEntries []*models.InstrumentEntry) []*models.Instrument {
	instruments := make([]*models.Instrument, len(instrumentEntries))
	for i, instrumentEntry := range instrumentEntries {
		instruments[i] = &models.Instrument{InstrumentEntry: instrumentEntry}
	}
	return instruments
}

// InstrumentsToMap transforms a slice of models.Instrument's into map, where the key is the ID and the value the models.Instrument.
func (i *InstrumentSvc) InstrumentsToMap(instruments []*models.Instrument) map[uint]*models.Instrument {
	instrumentMap := make(map[uint]*models.Instrument, len(instruments))
	for _, instrument := range instruments {
		instrumentMap[instrument.ID] = instrument
	}
	return instrumentMap
}
