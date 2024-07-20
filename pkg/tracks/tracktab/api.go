package tracktab

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing tracks of songs.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// LinkTrackToTab inserts a link between a track and a tab into the track_tab table.
func (a *API) LinkTrackToTab(trackId, tabId string) {
	_, err := a.Database.DB.Exec(insertTrackTabQuery, trackId, tabId)
	if err != nil {
		log.Printf("Failed linking track with ID '%s' and tab with ID '%s': %s", trackId, tabId, err.Error())
	} else {
		log.Printf("Successfully linked track with ID '%s' and tab with ID '%s'", trackId, tabId)
	}
}

// GetTabIDs retrieves the tab IDs for the provided track IDs.
func (a *API) GetTabIDs(trackID ...string) ([]string, error) {
	rows, err := a.Database.DB.Query(getTabIDs, trackID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tabIDs []string

	for rows.Next() {
		var tabID string
		err := rows.Scan(&tabID)
		if err != nil {
			return nil, err
		}
		tabIDs = append(tabIDs, tabID)
	}

	return tabIDs, nil
}
