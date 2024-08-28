package services

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
)

// InstrumentOps represents operations related to instruments.
type InstrumentOps interface {
	schemas.InstrumentSchema
	data.InstrumentData
	mappers.InstrumentMapper
	GetInstruments(instrumentID ...uint) ([]*models.Instrument, error)
}

// InstrumentSvc is responsible for managing instruments.
type InstrumentSvc struct {
	schemas.InstrumentSchema
	data.InstrumentData
	mappers.InstrumentMapper
}

// NewInstrumentSvc creates a new instance of InstrumentSvc.
func NewInstrumentSvc(schema schemas.InstrumentSchema, data data.InstrumentData, mapper mappers.InstrumentMapper) InstrumentOps {
	return &InstrumentSvc{
		InstrumentSchema: schema,
		InstrumentData:   data,
		InstrumentMapper: mapper,
	}
}

// GetInstruments retrieves instruments for the provided IDs.
func (i *InstrumentSvc) GetInstruments(instrumentID ...uint) ([]*models.Instrument, error) {
	instrumentEntries, err := i.GetInstrumentEntries(instrumentID...)
	if err != nil {
		return nil, err
	}

	instruments := i.InstrumentEntriesToInstruments(instrumentEntries)

	return instruments, nil
}
