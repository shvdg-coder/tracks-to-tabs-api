package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// InstrumentOps represents the operations related to instruments.
type InstrumentOps interface {
	data.InstrumentData
	mappers.InstrumentMapper
	GetInstruments(instrumentID ...uint) ([]*models.Instrument, error)
}

// InstrumentSvc is responsible for managing instruments.
type InstrumentSvc struct {
	data.InstrumentData
	mappers.InstrumentMapper
}

// NewInstrumentSvc creates a new instance of InstrumentSvc.
func NewInstrumentSvc(data data.InstrumentData, mapper mappers.InstrumentMapper) InstrumentOps {
	return &InstrumentSvc{InstrumentData: data, InstrumentMapper: mapper}
}

// GetInstruments retrieves instruments for the provided IDs.
func (s *InstrumentSvc) GetInstruments(instrumentID ...uint) ([]*models.Instrument, error) {
	return nil, nil
}
