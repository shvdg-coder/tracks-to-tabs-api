package services

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// DifficultyOps represents operations related to difficulties.
type DifficultyOps interface {
	data.DifficultyData
	mappers.DifficultyMapper
	GetDifficulties(difficultyID ...uint) ([]*models.Difficulty, error)
}

// DifficultySvc is responsible for managing difficulties.
type DifficultySvc struct {
	data.DifficultyData
	mappers.DifficultyMapper
}

// NewDifficultySvc instantiates a new instance of DifficultySvc.
func NewDifficultySvc(data data.DifficultyData, mapper mappers.DifficultyMapper) DifficultyOps {
	return &DifficultySvc{DifficultyData: data, DifficultyMapper: mapper}
}

// GetDifficulties retrieves the difficulties for the provided IDs.
func (d *DifficultySvc) GetDifficulties(difficultyID ...uint) ([]*models.Difficulty, error) {
	return nil, nil
}
