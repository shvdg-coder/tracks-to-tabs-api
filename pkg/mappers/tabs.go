package mappers

import (
	"github.com/google/uuid"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// MappingOperations represents operations related to tab data mapping.
type MappingOperations interface {
	TabsToMap(tabs []*diff.Tab) map[uuid.UUID]*diff.Tab
	MapInstrumentsToTabs(tabsMap map[uuid.UUID]*diff.Tab, instruments map[uint]*diff.InstrumentEntry) []*diff.Tab
	MapDifficultiesToTabs(tabsMap map[uuid.UUID]*diff.Tab, difficulties map[uint]*diff.DifficultyEntry) []*diff.Tab
	MapReferencesToTabs(tabsMap map[uuid.UUID]*diff.Tab, references []*diff.Reference) []*diff.Tab
}

// MappingService is responsible for mapping entities to tabs.
type MappingService struct {
	MappingOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService() MappingOperations {
	return &MappingService{}
}

// TabsToMap transforms a slice of tabs into a map where the key is the ID and the value the Tab.
func (m *MappingService) TabsToMap(tabs []*diff.Tab) map[uuid.UUID]*diff.Tab {
	tabsMap := make(map[uuid.UUID]*tabs.Tab)
	for _, tab := range tabs {
		tabsMap[tab.ID] = tab
	}
	return tabsMap
}

// MapInstrumentsToTabs todo:
func (m *MappingService) MapInstrumentsToTabs(tabsMap map[uuid.UUID]*diff.Tab, instruments map[uint]*diff.InstrumentEntry) []*diff.Tab {
	for _, tab := range tabsMap {
		instrument, ok := instruments[tab.Instrument.ID]
		if !ok {
			continue
		}
		tab.Instrument = instrument
	}
	var tabs []*diff.Tab
	for _, tab := range tabsMap {
		tabs = append(tabs, tab)
	}
	return tabs
}

// MapDifficultiesToTabs todo:
func (m *MappingService) MapDifficultiesToTabs(tabsMap map[uuid.UUID]*diff.Tab, difficulties map[uint]*diff.DifficultyEntry) []*diff.Tab {
	for _, tab := range tabsMap {
		difficulty, ok := difficulties[tab.Difficulty.ID]
		if !ok {
			continue
		}
		tab.Difficulty = difficulty
	}
	var tabs []*diff.Tab
	for _, tab := range tabsMap {
		tabs = append(tabs, tab)
	}
	return tabs
}

// MapReferencesToTabs maps references.Reference's to Tab's.
func (m *MappingService) MapReferencesToTabs(tabsMap map[uuid.UUID]*diff.Tab, references []*diff.Reference) []*diff.Tab {
	for _, reference := range references {
		tab, ok := tabsMap[reference.InternalID]
		if !ok {
			continue
		}
		tab.References = append(tab.References, reference)
	}
	var tabs []*diff.Tab
	for _, tab := range tabsMap {
		tabs = append(tabs, tab)
	}
	return tabs
}
