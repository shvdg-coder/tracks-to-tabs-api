package tabs

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing tabs.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// InsertTabs inserts multiple tabs in the tabs table.
func (a *API) InsertTabs(tabs ...*Tab) {
	for _, tab := range tabs {
		a.InsertTab(tab)
	}
}

// InsertTab inserts a new tab in the tabs table.
func (a *API) InsertTab(tab *Tab) {
	_, err := a.Database.DB.Exec(insertTabQuery, tab.ID, tab.Instrument.ID, tab.Difficulty.ID, tab.Description)
	if err != nil {
		log.Printf("Failed to insert tab with '%s', '%s' & Description: '%s': %s", tab.Instrument.Name, tab.Difficulty.Name, tab.Description, err.Error())
	} else {
		log.Printf("Successfully inserted tab with '%s', '%s' & Description: '%s'", tab.Instrument.Name, tab.Difficulty.Name, tab.Description)
	}
}

// GetTabs retrieves the tabs for the provided IDs.
func (a *API) GetTabs(tabID ...string) ([]*Tab, error) {
	return nil, nil
}
