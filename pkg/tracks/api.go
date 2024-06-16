package tracks

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

// CreateTracksTable creates the tracks table if it doesn't already exist.
func (a *API) CreateTracksTable() {
	_, err := a.Database.DB.Exec(createTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'tracks' table")
	}
}

// DropTracksTable drops the tracks table if it exists.
func (a *API) DropTracksTable() {
	_, err := a.Database.DB.Exec(dropTracksTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'tracks' table")
	}
}

// InsertTrack inserts a track into the tracks table.
func (a *API) InsertTrack(track *Track) {
	_, err := a.Database.DB.Exec(insertTrackQuery, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
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
