package tabs

import (
	"github.com/google/uuid"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	ins "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
)

// Operations represent all operations related to tabs.
type Operations interface {
	DataOperations
	MappingOperations
	GetTabsCascading(tabID ...uuid.UUID) ([]*Tab, error)
}

// Service is responsible for managing and retrieving tabs.
type Service struct {
	DataOperations
	MappingOperations
	InstrumentOps ins.Operations
	DifficultyOps diff.Operations
	ReferenceOps  ref.Operations
}

// NewService instantiates a Service.
func NewService(data DataOperations, mapping MappingOperations, instruments ins.Operations, difficulties diff.Operations, references ref.Operations) Operations {
	return &Service{
		DataOperations:    data,
		MappingOperations: mapping,
		InstrumentOps:     instruments,
		DifficultyOps:     difficulties,
		ReferenceOps:      references,
	}
}

// GetTabsCascading retrieves tabs, with entity references, for the provided IDs.
func (s *Service) GetTabsCascading(tabID ...uuid.UUID) ([]*Tab, error) {
	tabs, err := s.GetTabs(tabID...)
	if err != nil {
		return nil, err
	}

	references, err := s.ReferenceOps.GetReferences(tabID...)
	if err != nil {
		return nil, err
	}

	instrumentIDs, difficultyIDs := s.ExtractIDs(tabs)

	instruments, err := s.InstrumentOps.GetInstruments(instrumentIDs...)
	if err != nil {
		return nil, err
	}

	difficulties, err := s.DifficultyOps.GetDifficulties(difficultyIDs...)
	if err != nil {
		return nil, err
	}

	tabsMap := s.TabsToMap(tabs)
	instrumentsMap := s.InstrumentOps.InstrumentsToMap(instruments)
	difficultiesMap := s.DifficultyOps.DifficultiesToMap(difficulties)

	tabs = s.MapInstrumentsToTabs(tabsMap, instrumentsMap)
	tabs = s.MapDifficultiesToTabs(tabsMap, difficultiesMap)
	tabs = s.MapReferencesToTabs(tabsMap, references)

	return tabs, nil
}

// ExtractIDs extracts the instrument and difficulties IDs from tabs.
func (s *Service) ExtractIDs(tabs []*Tab) (instrumentIDs []uint, difficultyIDs []uint) {
	instrumentIDs = make([]uint, 0)
	difficultyIDs = make([]uint, 0)
	for _, tab := range tabs {
		instrumentIDs = append(instrumentIDs, tab.Instrument.ID)
		difficultyIDs = append(difficultyIDs, tab.Difficulty.ID)
	}
	return instrumentIDs, difficultyIDs
}
