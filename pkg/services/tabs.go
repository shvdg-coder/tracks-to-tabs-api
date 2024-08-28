package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
)

// TabOps represent all operations related to tabs.
type TabOps interface {
	schemas.TabSchema
	data.TabData
	mappers.TabMapper
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
	GetTabsCascading(tabID ...uuid.UUID) ([]*models.Tab, error)
	ExtractIDsFromTabs(tabs []*models.Tab) (tabIDs []uuid.UUID, instrumentIDs []uint, difficultyIDs []uint)
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

// GetTabs retrieves tabs, without entity references, for the provided IDs.
func (t *TabSvc) GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error) {
	tabEntries, err := t.GetTabEntries(tabID...)
	if err != nil {
		return nil, err
	}

	tabs := t.TabEntriesToTabs(tabEntries)

	return tabs, nil
}

// GetTabsCascading retrieves tabs, with entity references, for the provided IDs.
func (t *TabSvc) GetTabsCascading(tabID ...uuid.UUID) ([]*models.Tab, error) {
	tabs, err := t.GetTabs(tabID...)
	if err != nil {
		return nil, err
	}

	err = t.LoadInstruments(tabs...)
	if err != nil {
		return nil, err
	}

	err = t.LoadDifficulties(tabs...)
	if err != nil {
		return nil, err
	}

	err = t.LoadReferences(tabs...)
	if err != nil {
		return nil, err
	}

	return tabs, nil
}

// LoadInstruments loads the models.Instrument's for the given models.Tab's.
func (t *TabSvc) LoadInstruments(tabs ...*models.Tab) error {
	_, instrumentIDs, _ := t.ExtractIDsFromTabs(tabs)
	instruments, err := t.GetInstruments(instrumentIDs...)
	if err != nil {
		return err
	}

	tabsMap := t.TabsToMap(tabs)
	instrumentsMap := t.InstrumentsToMap(instruments)
	t.MapInstrumentsToTabs(tabsMap, instrumentsMap)

	return nil
}

// LoadDifficulties loads the models.Difficulty's for the given models.Tab's.
func (t *TabSvc) LoadDifficulties(tabs ...*models.Tab) error {
	_, _, difficultyIDs := t.ExtractIDsFromTabs(tabs)
	difficulties, err := t.GetDifficulties(difficultyIDs...)
	if err != nil {
		return err
	}

	tabsMap := t.TabsToMap(tabs)
	difficultiesMap := t.DifficultiesToMap(difficulties)
	t.MapDifficultiesToTabs(tabsMap, difficultiesMap)

	return nil
}

// LoadReferences loads the models.Reference's for the given models.Tab's.
func (t *TabSvc) LoadReferences(tabs ...*models.Tab) error {
	tabIDs, _, _ := t.ExtractIDsFromTabs(tabs)
	references, err := t.GetReferencesCascading(tabIDs...)
	if err != nil {
		return err
	}

	tabsMap := t.TabsToMap(tabs)
	t.MapReferencesToTabs(tabsMap, references)

	return nil
}

// ExtractIDsFromTabs retrieves the tab, instrument and difficulty ID's from the models.Tab's.
func (t *TabSvc) ExtractIDsFromTabs(tabs []*models.Tab) (tabIDs []uuid.UUID, instrumentIDs []uint, difficultyIDs []uint) {
	tabIDs = make([]uuid.UUID, len(tabs))
	instrumentIDs = make([]uint, len(tabs))
	difficultyIDs = make([]uint, len(tabs))

	for i, tab := range tabs {
		tabIDs[i] = tab.ID
		instrumentIDs[i] = tab.InstrumentID
		difficultyIDs[i] = tab.DifficultyID
	}

	return tabIDs, instrumentIDs, difficultyIDs
}
