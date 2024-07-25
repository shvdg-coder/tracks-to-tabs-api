package difficulties

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to difficulties in the database.
type DatabaseOperations interface {
	InsertDifficulties(difficulties ...*Difficulty)
	InsertDifficulty(difficulty *Difficulty)
	GetDifficulty(difficultyID string) (*Difficulty, error)
	GetDifficulties(difficultyID ...string) ([]*Difficulty, error)
}

// DatabaseService is for managing difficulties.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DatabaseOperations {
	return &DatabaseService{Database: database}
}

// InsertDifficulties inserts multiple difficulty levels.
func (a *DatabaseService) InsertDifficulties(difficulties ...*Difficulty) {
	for _, difficulty := range difficulties {
		a.InsertDifficulty(difficulty)
	}
}

// InsertDifficulty inserts a new difficulty level.
func (a *DatabaseService) InsertDifficulty(difficulty *Difficulty) {
	_, err := a.Database.DB.Exec(insertDifficultyQuery, difficulty.Name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with name: '%s': %s", difficulty.Name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with name: '%s'", difficulty.Name)
	}
}

// GetDifficulty retrieves a difficulty for the provided ID.
func (a *DatabaseService) GetDifficulty(difficultyID string) (*Difficulty, error) {
	difficulty, err := a.GetDifficulties(difficultyID)
	if err != nil {
		return nil, err
	}
	return difficulty[0], nil
}

// GetDifficulties retrieves difficulties for the provided IDs.
func (a *DatabaseService) GetDifficulties(difficultyID ...string) ([]*Difficulty, error) {
	rows, err := a.Database.DB.Query(getDifficultiesQuery, difficultyID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var difficulties []*Difficulty
	for rows.Next() {
		var difficulty Difficulty
		err := rows.Scan(&difficulty.ID, &difficulty.Name)
		if err != nil {
			return nil, err
		}
		difficulties = append(difficulties, &difficulty)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return difficulties, nil
}
