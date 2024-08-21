package difficulties

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to difficulties in the database.
type DataOperations interface {
	InsertDifficulties(difficulties ...*Difficulty)
	InsertDifficulty(difficulty *Difficulty)
	GetDifficulty(difficultyID uint) (*Difficulty, error)
	GetDifficulties(difficultyID ...uint) ([]*Difficulty, error)
}

// DataService is for managing difficulties.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertDifficulties inserts multiple difficulty levels.
func (d *DataService) InsertDifficulties(difficulties ...*Difficulty) {
	for _, difficulty := range difficulties {
		d.InsertDifficulty(difficulty)
	}
}

// InsertDifficulty inserts a new difficulty level.
func (d *DataService) InsertDifficulty(difficulty *Difficulty) {
	_, err := d.Exec(insertDifficultyQuery, difficulty.Name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with name: '%s': %s", difficulty.Name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with name: '%s'", difficulty.Name)
	}
}

// GetDifficulty retrieves a difficulty for the provided ID.
func (d *DataService) GetDifficulty(difficultyID uint) (*Difficulty, error) {
	difficulty, err := d.GetDifficulties(difficultyID)
	if err != nil {
		return nil, err
	}
	return difficulty[0], nil
}

// GetDifficulties retrieves difficulties for the provided IDs.
func (d *DataService) GetDifficulties(difficultyID ...uint) ([]*Difficulty, error) {
	rows, err := d.Query(getDifficultiesQuery, pq.Array(difficultyID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var difficulties []*Difficulty
	for rows.Next() {
		difficulty := &Difficulty{}
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
