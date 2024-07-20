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
