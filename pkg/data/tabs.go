package data

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TabData represents operations related to tabs in the database.
type TabData interface {
	InsertTabEntries(tabs ...*models.TabEntry) error
	GetTabEntries(tabID ...uuid.UUID) ([]*models.TabEntry, error)
}

// TabSvc is for managing queries.
type TabSvc struct {
	logic.DbOps
}

// NewTabSvc creates a new instance of the TabSvc struct.
func NewTabSvc(database logic.DbOps) TabData {
	return &TabSvc{DbOps: database}
}

// InsertTabEntries inserts multiple tabs in the tabs table.
func (d *TabSvc) InsertTabEntries(tabs ...*models.TabEntry) error {
	data := make([][]interface{}, len(tabs))

	for i, tab := range tabs {
		data[i] = tab.Fields()
	}

	fieldNames := []string{"id", "instrument_id", "difficulty_id", "description"}
	return d.BulkInsert("tabs", fieldNames, data)
}

// GetTabEntries retrieves tab entries, without entity references, for the provided IDs.
func (d *TabSvc) GetTabEntries(tabIDs ...uuid.UUID) ([]*models.TabEntry, error) {
	return logic.BatchGet(d, batchSize, queries.GetTabs, tabIDs, scanTabEntry)
}

// scanTabEntry scans a row into a models.TabEntry.
func scanTabEntry(rows *sql.Rows) (*models.TabEntry, error) {
	tab := &models.TabEntry{}
	if err := rows.Scan(&tab.ID, &tab.InstrumentID, &tab.DifficultyID, &tab.Description); err != nil {
		return nil, err
	}
	return tab, nil
}
