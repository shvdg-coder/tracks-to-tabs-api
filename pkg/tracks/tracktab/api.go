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

// GetTrackToTabLink retrieves the 'track to tab' link for the provided ID.
func (a *API) GetTrackToTabLink(trackID string) (*TrackTab, error) {
	trackTabLinks, err := a.GetTrackToTabLinks(trackID)
	if err != nil {
		return nil, err
	}
	return trackTabLinks[0], nil
}

// GetTrackToTabLinks retrieves the 'track to tab' links for the provided track IDs.
func (a *API) GetTrackToTabLinks(trackID ...string) ([]*TrackTab, error) {
	rows, err := a.Database.DB.Query(getTrackTabLinks, trackID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trackTabLinks []*TrackTab
	for rows.Next() {
		var trackTabLink *TrackTab
		err := rows.Scan(&trackTabLink.TrackID, &trackTabLink.TabID)
		if err != nil {
			return nil, err
		}
		trackTabLinks = append(trackTabLinks, trackTabLink)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return trackTabLinks, nil
}
