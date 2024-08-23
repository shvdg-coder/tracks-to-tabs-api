package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// TabData represents operations related to tabs in the database.
type TabData interface {
	InsertTabs(tabs ...*models.Tab)
	InsertTab(tab *models.Tab)
	GetTab(tabID uuid.UUID) (*models.Tab, error)
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
}

// TabSvc is for managing queries.
type TabSvc struct {
	logic.DbOperations
}

// NewTabSvc creates a new instance of the TabSvc struct.
func NewTabSvc(database logic.DbOperations) TabData {
	return &TabSvc{DbOperations: database}
}

// InsertTabs inserts multiple tabs in the tabs table.
func (d *TabSvc) InsertTabs(tabs ...*models.Tab) {
	for _, tab := range tabs {
		d.InsertTab(tab)
	}
}

// InsertTab inserts a new tab in the tabs table.
func (d *TabSvc) InsertTab(tab *models.Tab) {
	_, err := d.Exec(queries.InsertTab, tab.ID, tab.Instrument.ID, tab.Difficulty.ID, tab.Description)
	if err != nil {
		log.Printf("Failed to insert tab with '%s', '%s' & Description: '%s': %s", tab.Instrument.Name, tab.Difficulty.Name, tab.Description, err.Error())
	} else {
		log.Printf("Successfully inserted tab with '%s', '%s' & Description: '%s'", tab.Instrument.Name, tab.Difficulty.Name, tab.Description)
	}
}

// GetTab retrieves the tab, without entity references, for the provided tab ID.
func (d *TabSvc) GetTab(tabID uuid.UUID) (*models.Tab, error) {
	tabs, err := d.GetTabs(tabID)
	if err != nil {
		return nil, err
	}
	return tabs[0], nil
}

// GetTabs retrieves the tabs, without entity references, for the provided IDs.
func (d *TabSvc) GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error) {
	rows, err := d.Query(queries.GetTabs, pq.Array(tabID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabs []*models.Tab
	for rows.Next() {
		tab := &models.Tab{}
		instrument := &models.InstrumentEntry{}
		difficulty := &models.DifficultyEntry{}
		err := rows.Scan(&tab.ID, &instrument.ID, &difficulty.ID, &tab.Description)
		if err != nil {
			return nil, err
		}

		tab.Instrument = instrument
		tab.Difficulty = difficulty

		tabs = append(tabs, tab)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tabs, nil
}
