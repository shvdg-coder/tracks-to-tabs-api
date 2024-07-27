package tracktab

import "github.com/google/uuid"

// Operations represents all operations related to 'track to tab' links.
type Operations interface {
	DataOperations
	ExtractTabIDs(trackTabs []*TrackTab) []uuid.UUID
}

// Service is responsible for managing and retrieving 'track to tab' links.
type Service struct {
	DataOperations
}

// NewService instantiates a Service.
func NewService(database DataOperations) Operations {
	return &Service{DataOperations: database}
}

// ExtractTabIDs retrieves the tab IDs from each TrackTab.
func (service *Service) ExtractTabIDs(trackTabs []*TrackTab) []uuid.UUID {
	var tabIDs []uuid.UUID
	for _, link := range trackTabs {
		tabIDs = append(tabIDs, link.TabID)
	}
	return tabIDs
}
