package tracktab

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to 'track to tab' links.
type DatabaseOperations interface {
	LinkTrackToTab(trackID, tabID string)
	GetTrackToTabLink(trackID string) (*TrackTab, error)
	GetTrackToTabLinks(trackID ...string) ([]*TrackTab, error)
}

// DatabaseService is for managing tracks of songs.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) *DatabaseService {
	return &DatabaseService{Database: database}
}

// LinkTrackToTab inserts a link between a track and a tab into the track_tab table.
func (a *DatabaseService) LinkTrackToTab(trackId, tabId string) {
	_, err := a.Database.DB.Exec(insertTrackTabQuery, trackId, tabId)
	if err != nil {
		log.Printf("Failed linking track with ID '%s' and tab with ID '%s': %s", trackId, tabId, err.Error())
	} else {
		log.Printf("Successfully linked track with ID '%s' and tab with ID '%s'", trackId, tabId)
	}
}

// GetTrackToTabLink retrieves the 'track to tab' link for the provided ID.
func (a *DatabaseService) GetTrackToTabLink(trackID string) (*TrackTab, error) {
	trackTabLinks, err := a.GetTrackToTabLinks(trackID)
	if err != nil {
		return nil, err
	}
	return trackTabLinks[0], nil
}

// GetTrackToTabLinks retrieves the 'track to tab' links for the provided track IDs.
func (a *DatabaseService) GetTrackToTabLinks(trackID ...string) ([]*TrackTab, error) {
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
