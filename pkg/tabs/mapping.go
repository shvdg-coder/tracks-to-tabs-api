package tabs

import "github.com/google/uuid"

// MappingOperations represents operations related to tab data mapping.
type MappingOperations interface {
	TabsToMap(tabs []*Tab) map[uuid.UUID]*Tab
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
func (m *MappingService) TabsToMap(tabs []*Tab) map[uuid.UUID]*Tab {
	tabsMap := make(map[uuid.UUID]*Tab)
	for _, tab := range tabs {
		tabsMap[tab.ID] = tab
	}
	return tabsMap
}
