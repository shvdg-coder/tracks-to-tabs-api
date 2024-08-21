package tabs

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	ins "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"log"
)

// DataOperations represents operations related to tabs in the database.
type DataOperations interface {
	InsertTabs(tabs ...*Tab)
	InsertTab(tab *Tab)
	GetTab(tabID uuid.UUID) (*Tab, error)
	GetTabs(tabID ...uuid.UUID) ([]*Tab, error)
}

// DataService is for managing tabs.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertTabs inserts multiple tabs in the tabs table.
func (d *DataService) InsertTabs(tabs ...*Tab) {
	for _, tab := range tabs {
		d.InsertTab(tab)
	}
}

// InsertTab inserts a new tab in the tabs table.
func (d *DataService) InsertTab(tab *Tab) {
	_, err := d.Exec(insertTabQuery, tab.ID, tab.Instrument.ID, tab.Difficulty.ID, tab.Description)
	if err != nil {
		log.Printf("Failed to insert tab with '%s', '%s' & Description: '%s': %s", tab.Instrument.Name, tab.Difficulty.Name, tab.Description, err.Error())
	} else {
		log.Printf("Successfully inserted tab with '%s', '%s' & Description: '%s'", tab.Instrument.Name, tab.Difficulty.Name, tab.Description)
	}
}

// GetTab retrieves the tab, without entity references, for the provided tab ID.
func (d *DataService) GetTab(tabID uuid.UUID) (*Tab, error) {
	tabs, err := d.GetTabs(tabID)
	if err != nil {
		return nil, err
	}
	return tabs[0], nil
}

// GetTabs retrieves the tabs, without entity references, for the provided IDs.
func (d *DataService) GetTabs(tabID ...uuid.UUID) ([]*Tab, error) {
	rows, err := d.Query(getTabsQuery, pq.Array(tabID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabs []*Tab
	for rows.Next() {
		tab := &Tab{}
		instrument := &ins.Instrument{}
		difficulty := &diff.Difficulty{}
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
