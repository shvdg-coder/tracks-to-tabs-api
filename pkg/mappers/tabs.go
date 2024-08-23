package mappers

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// TabMapper represents operations related to tab data mapping.
type TabMapper interface {
	TabsToMap(tabs []*models.Tab) map[uuid.UUID]*models.Tab
	MapInstrumentsToTabs(tabsMap map[uuid.UUID]*models.Tab, instruments map[uint]*models.InstrumentEntry) map[uuid.UUID]*models.Tab
	MapDifficultiesToTabs(tabsMap map[uuid.UUID]*models.Tab, difficulties map[uint]*models.DifficultyEntry) map[uuid.UUID]*models.Tab
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

// TabsToMap transforms a slice of tabs into a map where the key is the ID and the value the Tab.
func (m *TabSvc) TabsToMap(tabs []*models.Tab) map[uuid.UUID]*models.Tab {
	tabsMap := make(map[uuid.UUID]*models.Tab)
	for _, tab := range tabs {
		tabsMap[tab.ID] = tab
	}
	return tabsMap
}

// MapInstrumentsToTabs maps models.Instrument's to models.Tab's.
func (m *TabSvc) MapInstrumentsToTabs(tabsMap map[uuid.UUID]*models.Tab, instruments map[uint]*models.InstrumentEntry) map[uuid.UUID]*models.Tab {
	for _, tab := range tabsMap {
		instrument := instruments[tab.Instrument.ID]
		tab.Instrument = instrument
	}
	return tabsMap
}

// MapDifficultiesToTabs maps models.Difficulty's to models.Tab's.
func (m *TabSvc) MapDifficultiesToTabs(tabsMap map[uuid.UUID]*models.Tab, difficulties map[uint]*models.DifficultyEntry) map[uuid.UUID]*models.Tab {
	for _, tab := range tabsMap {
		difficulty := difficulties[tab.Difficulty.ID]
		tab.Difficulty = difficulty
	}
	return tabsMap
}

// MapReferencesToTabs maps models.Reference's to models.Tab's.
func (m *TabSvc) MapReferencesToTabs(tabsMap map[uuid.UUID]*models.Tab, references []*models.Reference) map[uuid.UUID]*models.Tab {
	for _, reference := range references {
		tab := tabsMap[reference.InternalID]
		tab.References = append(tab.References, reference)
	}
	return tabsMap
}
