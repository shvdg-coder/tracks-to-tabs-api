package tabs

import "github.com/google/uuid"

// MappingOperations represents operations related to tab data mapping.
type MappingOperations interface {
	GetTabsAsMap(tabID ...uuid.UUID) (map[uuid.UUID]*Tab, error)
}

// MappingService is responsible for mapping entities to tabs.
type MappingService struct {
	DatabaseOperations
}

// NewMappingService creates a new instance of MappingService.
func NewMappingService(tabs DatabaseOperations) MappingOperations {
	return &MappingService{DatabaseOperations: tabs}
}

// GetTabsAsMap retrieves tabs for the provided IDs, and creates a map where the key is the ID and the value the Tab.
func (m *MappingService) GetTabsAsMap(tabID ...uuid.UUID) (map[uuid.UUID]*Tab, error) {
	tabs, err := m.GetTabs(tabID...)
	if err != nil {
		return nil, err
	}

	tabsMap := make(map[uuid.UUID]*Tab)
	for _, tab := range tabs {
		tabsMap[tab.ID] = tab
	}

	return tabsMap, nil
}
