package difficulties

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing difficulties.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// InsertDifficulties inserts multiple difficulty levels.
func (a *API) InsertDifficulties(difficulties ...*Difficulty) {
	for _, difficulty := range difficulties {
		a.InsertDifficulty(difficulty)
	}
}

// InsertDifficulty inserts a new difficulty level.
func (a *API) InsertDifficulty(difficulty *Difficulty) {
	_, err := a.Database.DB.Exec(insertDifficultyQuery, difficulty.Name)
	if err != nil {
		log.Printf("Failed inserting difficulty level with name: '%s': %s", difficulty.Name, err.Error())
	} else {
		log.Printf("Successfully inserted difficulty level with name: '%s'", difficulty.Name)
	}
}

// GetDifficulty retrieves a difficulty for the provided ID.
func (a *API) GetDifficulty(difficultyID string) (*Difficulty, error) {
	difficulty, err := a.GetDifficulties(difficultyID)
	if err != nil {
		return nil, err
	}
	return difficulty[0], nil
}

// GetDifficulties retrieves difficulties for the provided IDs.
func (a *API) GetDifficulties(difficultyID ...string) ([]*Difficulty, error) {
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
