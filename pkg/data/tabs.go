package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TabData represents operations related to tabs in the database.
type TabData interface {
	InsertTabEntries(tabs ...*models.TabEntry) error
	GetTabEntries(tabID ...uuid.UUID) ([]*models.TabEntry, error)
	GetTabEntry(tabID uuid.UUID) (*models.TabEntry, error)
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

// GetTabEntry retrieves a tab entry, without entity references, for the provided tab ID.
func (d *TabSvc) GetTabEntry(tabID uuid.UUID) (*models.TabEntry, error) {
	tabs, err := d.GetTabEntries(tabID)
	if err != nil {
		return nil, err
	}
	return tabs[0], nil
}

// GetTabEntries retrieves tab entries, without entity references, for the provided IDs.
func (d *TabSvc) GetTabEntries(tabID ...uuid.UUID) ([]*models.TabEntry, error) {
	rows, err := d.DB().Query(queries.GetTabs, pq.Array(tabID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabs []*models.TabEntry
	for rows.Next() {
		tab := &models.TabEntry{}
		err := rows.Scan(&tab.ID, &tab.InstrumentID, &tab.DifficultyID, &tab.Description)
		if err != nil {
			return nil, err
		}
		tabs = append(tabs, tab)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tabs, nil
}
