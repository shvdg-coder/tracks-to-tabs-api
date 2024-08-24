package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TabMapper represents operations related to tab data mapping.
type TabMapper interface {
	TabEntriesToTabs(tabEntries []*models.TabEntry) []*models.Tab
	TabsToMap(tabs []*models.Tab) map[uuid.UUID]*models.Tab
	MapToTabs(tabsMap map[uuid.UUID]*models.Tab) []*models.Tab
	MapInstrumentsToTabs(tabsMap map[uuid.UUID]*models.Tab, instruments map[uint]*models.Instrument) map[uuid.UUID]*models.Tab
	MapDifficultiesToTabs(tabsMap map[uuid.UUID]*models.Tab, difficulties map[uint]*models.Difficulty) map[uuid.UUID]*models.Tab
	MapReferencesToTabs(tabsMap map[uuid.UUID]*models.Tab, references []*models.Reference) map[uuid.UUID]*models.Tab
}

// TabSvc is responsible for mapping entities to tabs.
type TabSvc struct {
	TabMapper
}

// NewTabSvc creates a new instance of ReferenceSvc.
func NewTabSvc() TabMapper {
	return &TabSvc{}
}

// TabEntriesToTabs transforms the models.TabEntry's to models.Tab's.
func (t *TabSvc) TabEntriesToTabs(tabEntries []*models.TabEntry) []*models.Tab {
	tabs := make([]*models.Tab, len(tabEntries))
	for i, tabEntry := range tabEntries {
		tabs[i] = &models.Tab{
			TabEntry:   tabEntry,
			Instrument: &models.Instrument{InstrumentEntry: &models.InstrumentEntry{ID: tabEntry.InstrumentID}},
			Difficulty: &models.Difficulty{DifficultyEntry: &models.DifficultyEntry{ID: tabEntry.DifficultyID}},
		}
	}
	return tabs
}

// TabsToMap transforms a slice of tabs into a map where the key is the ID and the value the Tab.
func (t *TabSvc) TabsToMap(tabs []*models.Tab) map[uuid.UUID]*models.Tab {
	tabsMap := make(map[uuid.UUID]*models.Tab)
	for _, tab := range tabs {
		tabsMap[tab.ID] = tab
	}
	return tabsMap
}

// MapToTabs transforms a map of models.Tab's into a slice of models.Tab's.
func (t *TabSvc) MapToTabs(tabsMap map[uuid.UUID]*models.Tab) []*models.Tab {
	tabs := make([]*models.Tab, len(tabsMap))
	for _, tab := range tabsMap {
		tabs = append(tabs, tab)
	}
	return tabs
}

// MapInstrumentsToTabs maps models.Instrument's to models.Tab's, by updating the provided models.Tab's map and returning it.
func (t *TabSvc) MapInstrumentsToTabs(tabsMap map[uuid.UUID]*models.Tab, instruments map[uint]*models.Instrument) map[uuid.UUID]*models.Tab {
	for _, tab := range tabsMap {
		instrument := instruments[tab.Instrument.ID]
		tab.Instrument = instrument
	}
	return tabsMap
}

// MapDifficultiesToTabs maps models.Difficulty's to models.Tab's, by updating the provided models.Tab's map and returning it.
func (t *TabSvc) MapDifficultiesToTabs(tabsMap map[uuid.UUID]*models.Tab, difficulties map[uint]*models.Difficulty) map[uuid.UUID]*models.Tab {
	for _, tab := range tabsMap {
		difficulty := difficulties[tab.Difficulty.ID]
		tab.Difficulty = difficulty
	}
	return tabsMap
}

// MapReferencesToTabs maps models.Reference's to models.Tab's.
func (t *TabSvc) MapReferencesToTabs(tabsMap map[uuid.UUID]*models.Tab, references []*models.Reference) map[uuid.UUID]*models.Tab {
	for _, reference := range references {
		tab := tabsMap[reference.InternalID]
		tab.References = append(tab.References, reference)
	}
	return tabsMap
}
