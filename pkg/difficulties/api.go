package difficulties

import (
	"database/sql"
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

// CreateDifficultiesTable creates a difficulties table.
func (a *API) CreateDifficultiesTable() {
	_, err := a.Database.DB.Exec(createDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'difficulties' table.")
	}
}

// DropDifficultiesTable drops the difficulties table.
func (a *API) DropDifficultiesTable() {
	_, err := a.Database.DB.Exec(dropDifficultiesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'difficulties' table.")
	}
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

// GetDifficulties retrieves the difficulties.
func (a *API) GetDifficulties() []*Difficulty {
	rows, err := a.Database.DB.Query(getDifficultiesQuery)
	if err != nil {
		log.Printf("Failed to get difficulties: %s", err)
		return nil
	}

	difficulties := rowsToDifficulties(rows)
	defer closeRows(rows)

	return difficulties
}

// rowsToDifficulties converts the given *sql.Rows into a slice of *Difficulty objects.
func rowsToDifficulties(rows *sql.Rows) []*Difficulty {
	var difficulties []*Difficulty
	for rows.Next() {
		difficulty := rowsToDifficulty(rows)
		if difficulty != nil {
			difficulties = append(difficulties, difficulty)
		}
	}
	return difficulties
}

// rowsToDifficulty scans the SQL row into a Difficulty struct.
func rowsToDifficulty(rows *sql.Rows) *Difficulty {
	var difficulty Difficulty
	err := rows.Scan(&difficulty.ID, &difficulty.Name)
	if err != nil {
		log.Printf("Unable to scan difficulty: %s", err.Error())
		return nil
	}
	return &difficulty
}

// closeRows closes the SQL rows and logs error if any.
func closeRows(rows *sql.Rows) {
	err := rows.Err()
	if err != nil {
		log.Printf("Error while processing rows: %s", err.Error())
	}
	err = rows.Close()
	if err != nil {
		log.Printf("Failed to close rows: %s", err.Error())
	}
}
