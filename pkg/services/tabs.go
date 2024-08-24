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
	tabEntries, err := t.GetTabEntries(tabID...)
	if err != nil {
		return nil, err
	}

	instrumentIDs, difficultyIDs := t.ExtractEntityIDs(tabEntries)

	instruments, err := t.GetInstruments(instrumentIDs...)
	if err != nil {
		return nil, err
	}

	difficulties, err := t.GetDifficulties(difficultyIDs...)
	if err != nil {
		return nil, err
	}

	references, err := t.GetReferences(tabID...)
	if err != nil {
		return nil, err
	}

	tabs := t.TabEntriesToTabs(tabEntries)
	tabsMap := t.TabsToMap(tabs)
	instrumentsMap := t.InstrumentsToMap(instruments)
	difficultiesMap := t.DifficultiesToMap(difficulties)

	tabsMap = t.MapInstrumentsToTabs(tabsMap, instrumentsMap)
	tabsMap = t.MapDifficultiesToTabs(tabsMap, difficultiesMap)
	tabsMap = t.MapReferencesToTabs(tabsMap, references)
	tabs = t.MapToTabs(tabsMap)

	return tabs, nil
}

// ExtractEntityIDs extracts the instrument and difficulties IDs from models.TabEntry's.
func (t *TabSvc) ExtractEntityIDs(tabs []*models.TabEntry) (instrumentIDs []uint, difficultyIDs []uint) {
	instrumentIDs = make([]uint, len(tabs))
	difficultyIDs = make([]uint, len(tabs))
	for _, tab := range tabs {
		instrumentIDs = append(instrumentIDs, tab.InstrumentID)
		difficultyIDs = append(difficultyIDs, tab.DifficultyID)
	}
	return instrumentIDs, difficultyIDs
}
