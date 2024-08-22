package database

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// DifficultyOps represents operations related to difficulties in the database.
type DifficultyOps interface {
	InsertDifficulties(difficulties ...*models.DifficultyEntry)
	InsertDifficulty(difficulty *models.DifficultyEntry)
	GetDifficulty(difficultyID uint) (*models.DifficultyEntry, error)
	GetDifficulties(difficultyID ...uint) ([]*models.DifficultyEntry, error)
}

// DifficultySvc is for managing difficulties.
type DifficultySvc struct {
	logic.DbOperations
}

// NewDifficultySvc creates a new instance of the DifficultySvc struct.
func NewDifficultySvc(database logic.DbOperations) DifficultyOps {
	return &DifficultySvc{DbOperations: database}
}

// InsertDifficulties inserts multiple difficulty levels.
func (d *DifficultySvc) InsertDifficulties(difficulties ...*models.DifficultyEntry) {
	for _, difficulty := range difficulties {
		d.InsertDifficulty(difficulty)
	}
}

// InsertDifficulty inserts a new difficulty level.
func (d *DifficultySvc) InsertDifficulty(difficulty *models.DifficultyEntry) {
	_, err := d.Exec(queries.InsertDifficulty, difficulty.Name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with name: '%s': %s", difficulty.Name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with name: '%s'", difficulty.Name)
	}
}

// GetDifficulty retrieves a difficulty for the provided ID.
func (d *DifficultySvc) GetDifficulty(difficultyID uint) (*models.DifficultyEntry, error) {
	difficulty, err := d.GetDifficulties(difficultyID)
	if err != nil {
		return nil, err
	}
	return difficulty[0], nil
}

// GetDifficulties retrieves difficulties for the provided IDs.
func (d *DifficultySvc) GetDifficulties(difficultyID ...uint) ([]*models.DifficultyEntry, error) {
	rows, err := d.Query(queries.GetDifficulties, pq.Array(difficultyID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var difficulties []*models.DifficultyEntry
	for rows.Next() {
		difficulty := &models.DifficultyEntry{}
		err := rows.Scan(&difficulty.ID, &difficulty.Name)
		if err != nil {
			return nil, err
		}
		difficulties = append(difficulties, difficulty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return difficulties, nil
}
