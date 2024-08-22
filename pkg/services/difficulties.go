package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// ArtistTrackOps represents operations related to difficulties.
type Operations interface {
	database.TabsOps
	DifficultiesToMap(difficulties []*models.DifficultyEntry) map[uint]*models.DifficultyEntry
}

// ArtistTrackSvc is responsible for managing difficulties.
type Service struct {
	database.TabsOps
}

// NewTrackSvc instantiates a new instance of ArtistTrackSvc.
func NewService(data database.TabsOps) Operations {
	return &Service{TabsOps: data}
}

// DifficultiesToMap transforms a slice of DifficultyEntry's into a map, where the key is the ID and the value the DifficultyEntry.
func (s *Service) DifficultiesToMap(difficulties []*models.DifficultyEntry) map[uint]*models.DifficultyEntry {
	difficultyMap := make(map[uint]*models.DifficultyEntry)
	for _, difficulty := range difficulties {
		difficultyMap[difficulty.ID] = difficulty
	}
	return difficultyMap
}
