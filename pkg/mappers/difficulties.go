package mappers

import "github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"

// DifficultyMapper represents operations related to difficulty data mapping.
type DifficultyMapper interface {
	DifficultyEntriesToDifficulties(difficultyEntries []*models.DifficultyEntry) []*models.Difficulty
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

// DifficultyEntriesToDifficulties transforms a slice of models.DifficultyEntry's into a slice of models.Difficulty's.
func (d *DifficultySvc) DifficultyEntriesToDifficulties(difficultyEntries []*models.DifficultyEntry) []*models.Difficulty {
	difficulties := make([]*models.Difficulty, len(difficultyEntries))
	for i, entry := range difficultyEntries {
		difficulties[i] = &models.Difficulty{DifficultyEntry: entry}
	}
	return difficulties
}
