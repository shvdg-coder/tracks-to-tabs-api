package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TabOps represent all operations related to tabs.
type TabOps interface {
	data.TabData
	mappers.TabMapper
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
}

// TabSvc is responsible for managing and retrieving tabs.
type TabSvc struct {
	data.TabData
	mappers.TabMapper
	InstrumentOps
	DifficultyOps
	ReferenceOps
}

// NewTabSvc instantiates a TabSvc.
func NewTabSvc(data data.TabData, mapper mappers.TabMapper, instruments InstrumentOps, difficulties DifficultyOps, references ReferenceOps) TabOps {
	return &TabSvc{
		TabData:       data,
		TabMapper:     mapper,
		InstrumentOps: instruments,
		DifficultyOps: difficulties,
		ReferenceOps:  references,
	}
}

// GetTabs retrieves tabs, with entity references, for the provided IDs.
func (t *TabSvc) GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error) {
	return nil, nil
}

// ExtractIDs extracts the instrument and difficulties IDs from tabs.
func (t *TabSvc) ExtractIDs(tabs []*models.Tab) (instrumentIDs []uint, difficultyIDs []uint) {
	instrumentIDs = make([]uint, 0)
	difficultyIDs = make([]uint, 0)
	for _, tab := range tabs {
		instrumentIDs = append(instrumentIDs, tab.Instrument.ID)
		difficultyIDs = append(difficultyIDs, tab.Difficulty.ID)
	}
	return instrumentIDs, difficultyIDs
}
