package services

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
)

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
)

// DifficultyOps combines interfaces representing all operations related to difficulties.
type DifficultyOps interface {
	schemas.DifficultySchema
	data.DifficultyData
	mappers.DifficultyMapper
	GetDifficulties(difficultyID ...uint) ([]*models.Difficulty, error)
}

// DifficultySvc is responsible for managing difficulties.
type DifficultySvc struct {
	schemas.DifficultySchema
	data.DifficultyData
	mappers.DifficultyMapper
}

// NewDifficultySvc instantiates a new instance of DifficultySvc.
func NewDifficultySvc(schema schemas.DifficultySchema, data data.DifficultyData, mapper mappers.DifficultyMapper) DifficultyOps {
	return &DifficultySvc{
		DifficultySchema: schema,
		DifficultyData:   data,
		DifficultyMapper: mapper,
	}
}

// GetDifficulties retrieves the difficulties for the provided IDs.
func (d *DifficultySvc) GetDifficulties(difficultyID ...uint) ([]*models.Difficulty, error) {
	difficultyEntries, err := d.GetDifficultyEntries(difficultyID...)
	if err != nil {
		return nil, err
	}

	difficulties := d.DifficultyEntriesToDifficulties(difficultyEntries)

	return difficulties, nil
}
