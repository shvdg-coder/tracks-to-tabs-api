package data

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// DifficultyData represents operations related to difficulties in the database.
type DifficultyData interface {
	InsertDifficultyEntries(difficulties ...*models.DifficultyEntry) error
	GetDifficultyEntries(difficultyID ...uint) ([]*models.DifficultyEntry, error)
	GetDifficultyEntry(difficultyID uint) (*models.DifficultyEntry, error)
}

// DifficultySvc is for managing difficulties.
type DifficultySvc struct {
	logic.DbOps
}

// NewDifficultySvc creates a new instance of the DifficultySvc struct.
func NewDifficultySvc(database logic.DbOps) DifficultyData {
	return &DifficultySvc{DbOps: database}
}

// InsertDifficultyEntries inserts multiple difficulty levels.
func (d *DifficultySvc) InsertDifficultyEntries(difficulties ...*models.DifficultyEntry) error {
	data := make([][]interface{}, len(difficulties))

	for i, difficulty := range difficulties {
		data[i] = logic.GetFields(difficulty)
	}

	return d.BulkInsert("difficulties", logic.GetFieldNames("db", &models.DifficultyEntry{}), data)
}

// GetDifficultyEntry retrieves a difficulty entry, without the entity references, for the provided ID.
func (d *DifficultySvc) GetDifficultyEntry(difficultyID uint) (*models.DifficultyEntry, error) {
	difficulty, err := d.GetDifficultyEntries(difficultyID)
	if err != nil {
		return nil, err
	}
	return difficulty[0], nil
}

// GetDifficultyEntries retrieves difficulty entries, without entity references, for the provided IDs.
func (d *DifficultySvc) GetDifficultyEntries(difficultyID ...uint) ([]*models.DifficultyEntry, error) {
	rows, err := d.DB().Query(queries.GetDifficulties, pq.Array(difficultyID))
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
