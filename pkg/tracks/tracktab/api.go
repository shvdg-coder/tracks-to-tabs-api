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

// CreateTrackTabTable creates a track_tab table if it doesn't already exist.
func (a *API) CreateTrackTabTable() {
	_, err := a.Database.DB.Exec(createTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'track_tab' table")
	}
}

// DropTrackTabTable drops the track_tab table if it exists.
func (a *API) DropTrackTabTable() {
	_, err := a.Database.DB.Exec(dropTrackTabTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'track_tab' table")
	}
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
