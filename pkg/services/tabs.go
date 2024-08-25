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
	GetTabsCascading(tabID ...uuid.UUID) ([]*models.Tab, error)
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
	instrumentIDs := t.ExtractInstrumentIDs(tabs)
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
	difficultyIDs := t.ExtractDifficultyIDs(tabs)
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
	references, err := t.GetReferences(t.ExtractTabIDs(tabs)...)
	if err != nil {
		return err
	}

	tabsMap := t.TabsToMap(tabs)
	t.MapReferencesToTabs(tabsMap, references)

	return nil
}

// ExtractTabIDs retrieves the ID's from the models.Tab's.
func (t *TabSvc) ExtractTabIDs(tabs []*models.Tab) []uuid.UUID {
	tabIDs := make([]uuid.UUID, len(tabs))
	for i, tab := range tabs {
		tabIDs[i] = tab.ID
	}
	return tabIDs
}

// ExtractInstrumentIDs retrieves the Instrument IDs from the models.Tab's.
func (t *TabSvc) ExtractInstrumentIDs(tabs []*models.Tab) []uint {
	instrumentIDs := make([]uint, len(tabs))
	for i, tab := range tabs {
		instrumentIDs[i] = tab.InstrumentID
	}
	return instrumentIDs
}

// ExtractDifficultyIDs retrieves the Difficulty IDs from the models.Tab's.
func (t *TabSvc) ExtractDifficultyIDs(tabs []*models.Tab) []uint {
	difficultyIDs := make([]uint, len(tabs))
	for i, tab := range tabs {
		difficultyIDs[i] = tab.DifficultyID
	}
	return difficultyIDs
}
