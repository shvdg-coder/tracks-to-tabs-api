package mappers

import "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"

// DifficultyMapper represents operations related to difficulty data mapping.
type DifficultyMapper interface {
	DifficultiesToMap(difficulties []*models.Difficulty) map[uint]*models.Difficulty
}

// DifficultySvc is responsible for mapping entities to difficulties.
type DifficultySvc struct {
	DifficultyMapper
}

// NewDifficultySvc creates a new instance of ReferenceSvc.
func NewDifficultySvc() DifficultyMapper {
	return &DifficultySvc{}
}

// DifficultiesToMap transforms a slice of models.Difficulty's into a map, where the key is the ID and the value the models.Difficulty.
func (d *DifficultySvc) DifficultiesToMap(difficulties []*models.Difficulty) map[uint]*models.Difficulty {
	difficultyMap := make(map[uint]*models.Difficulty, len(difficulties))
	for _, difficulty := range difficulties {
		difficultyMap[difficulty.ID] = difficulty
	}
	return difficultyMap
}
