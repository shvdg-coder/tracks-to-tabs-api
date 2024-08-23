package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
)

// InstrumentOps represents the operations related to instruments.
type InstrumentOps interface {
	data.InstrumentData
	mappers.InstrumentMapper
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
