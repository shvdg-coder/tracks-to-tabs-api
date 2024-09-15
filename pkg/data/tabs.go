package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TabData represents operations related to tabs in the database.
type TabData interface {
	InsertTabEntries(tabs ...*models.TabEntry)
	InsertTabEntry(tab *models.TabEntry)
	GetTabEntries(tabID ...uuid.UUID) ([]*models.TabEntry, error)
	GetTabEntry(tabID uuid.UUID) (*models.TabEntry, error)
}

// TabSvc is for managing queries.
type TabSvc struct {
	logic.DbOperations
}

// NewTabSvc creates a new instance of the TabSvc struct.
func NewTabSvc(database logic.DbOperations) TabData {
	return &TabSvc{DbOperations: database}
}

// InsertTabEntries inserts multiple tabs in the tabs table.
func (d *TabSvc) InsertTabEntries(tabs ...*models.TabEntry) {
	for _, tab := range tabs {
		d.InsertTabEntry(tab)
	}
}

// InsertTabEntry inserts a new tab in the tabs table.
func (d *TabSvc) InsertTabEntry(tab *models.TabEntry) {
	_, err := d.Exec(queries.InsertTab, tab.ID, tab.InstrumentID, tab.DifficultyID, tab.Description)
	if err != nil {
		log.Printf("Failed to insert tab: %s", err.Error())
	}
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
	rows, err := d.Query(queries.GetTabs, pq.Array(tabID))
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
