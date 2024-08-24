package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/schemas"
)

// TabOps represent all operations related to tabs.
type TabOps interface {
	schemas.TabSchema
	data.TabData
	mappers.TabMapper
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
}

// TabSvc is responsible for managing and retrieving tabs.
type TabSvc struct {
	schemas.TabSchema
	data.TabData
	mappers.TabMapper
	InstrumentOps
	DifficultyOps
	ReferenceOps
}

// NewTabSvc instantiates a TabSvc.
func NewTabSvc(schema schemas.TabSchema, data data.TabData, mapper mappers.TabMapper, instruments InstrumentOps, difficulties DifficultyOps, references ReferenceOps) TabOps {
	return &TabSvc{
		TabSchema:     schema,
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
func (t *TabSvc) ExtractEntityIDs(tabEntries []*models.TabEntry) (instrumentIDs []uint, difficultyIDs []uint) {
	instrumentIDMap, difficultyIDMap := make(map[uint]bool), make(map[uint]bool)
	for _, tabEntry := range tabEntries {
		instrumentIDMap[tabEntry.InstrumentID] = true
		difficultyIDMap[tabEntry.DifficultyID] = true
	}

	keysToSlice := func(inputMap map[uint]bool) []uint {
		outputSlice := make([]uint, 0)
		for key, _ := range inputMap {
			outputSlice = append(outputSlice, key)
		}
		return outputSlice
	}

	instrumentIDs = keysToSlice(instrumentIDMap)
	difficultyIDs = keysToSlice(difficultyIDMap)

	return instrumentIDs, difficultyIDs
}
