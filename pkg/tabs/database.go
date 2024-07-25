package tabs

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to tabs in the database.
type DatabaseOperations interface {
	InsertTabs(tabs ...*Tab)
	InsertTab(tab *Tab)
	GetTab(tabID string) (*Tab, error)
	GetTabs(tabID ...string) ([]*Tab, error)
}

// DatabaseService is for managing tabs.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DatabaseOperations {
	return &DatabaseService{Database: database}
}

// InsertTabs inserts multiple tabs in the tabs table.
func (a *DatabaseService) InsertTabs(tabs ...*Tab) {
	for _, tab := range tabs {
		a.InsertTab(tab)
	}
}

// InsertTab inserts a new tab in the tabs table.
func (a *DatabaseService) InsertTab(tab *Tab) {
	_, err := a.Database.DB.Exec(insertTabQuery, tab.ID, tab.Instrument.ID, tab.Difficulty.ID, tab.Description)
	if err != nil {
		log.Printf("Failed to insert tab with '%s', '%s' & Description: '%s': %s", tab.Instrument.Name, tab.Difficulty.Name, tab.Description, err.Error())
	} else {
		log.Printf("Successfully inserted tab with '%s', '%s' & Description: '%s'", tab.Instrument.Name, tab.Difficulty.Name, tab.Description)
	}
}

// GetTab retrieves the tab, without entity references, for the provided tab ID.
func (a *DatabaseService) GetTab(tabID string) (*Tab, error) {
	tabs, err := a.GetTabs(tabID)
	if err != nil {
		return nil, err
	}
	return tabs[0], nil
}

// GetTabs retrieves the tabs, without entity references, for the provided IDs.
func (a *DatabaseService) GetTabs(tabID ...string) ([]*Tab, error) {
	rows, err := a.Database.DB.Query(getTabsQuery, tabID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabs []*Tab
	for rows.Next() {
		var tab *Tab
		var instrumentID, difficultyID string
		err := rows.Scan(&tab.ID, instrumentID, difficultyID, &tab.Description)
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
