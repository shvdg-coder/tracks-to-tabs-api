package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents the operations related to instruments.
type Operations interface {
	database.TabsOps
	InstrumentsToMap(instruments []*models.InstrumentEntry) map[uint]*models.InstrumentEntry
}

// ArtistTrackSvc is responsible for managing instruments.
type Service struct {
	database.TabsOps
}

// NewTrackSvc creates a new instance of ArtistTrackSvc.
func NewService(data database.TabsOps) Operations {
	return &Service{TabsOps: data}
}

// InstrumentsToMap transforms a slice of InstrumentEntry's into map, where the key is the ID and the value the InstrumentEntry.
func (s *Service) InstrumentsToMap(instruments []*models.InstrumentEntry) map[uint]*models.InstrumentEntry {
	instrumentMap := make(map[uint]*models.InstrumentEntry)
	for _, instrument := range instruments {
		instrumentMap[instrument.ID] = instrument
	}
	return instrumentMap
}
