package tracktab

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to 'track to tab' links.
type DataOperations interface {
	LinkTrackToTab(trackID, tabID uuid.UUID)
	GetTrackToTabLink(trackID uuid.UUID) (*TrackTab, error)
	GetTrackToTabLinks(trackID ...uuid.UUID) ([]*TrackTab, error)
}

// DataService is for managing tracks of songs.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) *DataService {
	return &DataService{DbOperations: database}
}

// LinkTrackToTab inserts a link between a track and a tab into the track_tab table.
func (d *DataService) LinkTrackToTab(trackId, tabId uuid.UUID) {
	_, err := d.Exec(insertTrackTabQuery, trackId, tabId)
	if err != nil {
		log.Printf("Failed linking track with ID '%s' and tab with ID '%s': %s", trackId, tabId, err.Error())
	} else {
		log.Printf("Successfully linked track with ID '%s' and tab with ID '%s'", trackId, tabId)
	}
}

// GetTrackToTabLink retrieves the 'track to tab' link for the provided ID.
func (d *DataService) GetTrackToTabLink(trackID uuid.UUID) (*TrackTab, error) {
	trackTabLinks, err := d.GetTrackToTabLinks(trackID)
	if err != nil {
		return nil, err
	}
	return trackTabLinks[0], nil
}

// GetTrackToTabLinks retrieves the 'track to tab' links for the provided track IDs.
func (d *DataService) GetTrackToTabLinks(trackID ...uuid.UUID) ([]*TrackTab, error) {
	rows, err := d.Query(getTrackTabLinks, pq.Array(trackID))
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
