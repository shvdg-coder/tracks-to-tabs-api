package services

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represent all operations related to tabs.
type Operations interface {
	database.TabsOps
	mappers.MappingOperations
	GetTabsCascading(tabID ...uuid.UUID) ([]*models.Tab, error)
}

// ArtistTrackSvc is responsible for managing and retrieving tabs.
type Service struct {
	database.TabsOps
	mappers.MappingOperations
	InstrumentOps Operations
	DifficultyOps Operations
	ReferenceOps  Operations
}

// NewTrackSvc instantiates a ArtistTrackSvc.
func NewService(data database.TabsOps, mapping mappers.MappingOperations, instruments Operations, difficulties Operations, references Operations) Operations {
	return &Service{
		TabsOps:           data,
		MappingOperations: mapping,
		InstrumentOps:     instruments,
		DifficultyOps:     difficulties,
		ReferenceOps:      references,
	}
}

// GetTabsCascading retrieves tabs, with entity references, for the provided IDs.
func (s *Service) GetTabsCascading(tabID ...uuid.UUID) ([]*models.Tab, error) {
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
func (s *Service) ExtractIDs(tabs []*models.Tab) (instrumentIDs []uint, difficultyIDs []uint) {
	instrumentIDs = make([]uint, 0)
	difficultyIDs = make([]uint, 0)
	for _, tab := range tabs {
		instrumentIDs = append(instrumentIDs, tab.Instrument.ID)
		difficultyIDs = append(difficultyIDs, tab.Difficulty.ID)
	}
	return instrumentIDs, difficultyIDs
}
