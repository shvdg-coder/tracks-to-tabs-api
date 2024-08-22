package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// TrackTabOps represents operations related to 'track to tab' links.
type TrackTabOps interface {
	LinkTrackToTab(trackID, tabID uuid.UUID)
	GetTrackToTabLink(trackID uuid.UUID) (*models.TrackTabEntry, error)
	GetTrackToTabLinks(trackID ...uuid.UUID) ([]*models.TrackTabEntry, error)
}

// TrackTabSvc is for managing tracks of songs.
type TrackTabSvc struct {
	logic.DbOperations
}

// NewTrackTabSvc creates a new instance of the TrackTabSvc struct.
func NewTrackTabSvc(database logic.DbOperations) *TrackTabSvc {
	return &TrackTabSvc{DbOperations: database}
}

// LinkTrackToTab inserts a link between a track and a tab into the track_tab table.
func (d *TrackTabSvc) LinkTrackToTab(trackId, tabId uuid.UUID) {
	_, err := d.Exec(queries.InsertTrackTab, trackId, tabId)
	if err != nil {
		log.Printf("Failed linking track with ID '%s' and tab with ID '%s': %s", trackId, tabId, err.Error())
	} else {
		log.Printf("Successfully linked track with ID '%s' and tab with ID '%s'", trackId, tabId)
	}
}

// GetTrackToTabLink retrieves the 'track to tab' link for the provided ID.
func (d *TrackTabSvc) GetTrackToTabLink(trackID uuid.UUID) (*models.TrackTabEntry, error) {
	trackTabLinks, err := d.GetTrackToTabLinks(trackID)
	if err != nil {
		return nil, err
	}
	return trackTabLinks[0], nil
}

// GetTrackToTabLinks retrieves the 'track to tab' links for the provided track IDs.
func (d *TrackTabSvc) GetTrackToTabLinks(trackID ...uuid.UUID) ([]*models.TrackTabEntry, error) {
	rows, err := d.Query(queries.GetTrackTabLinks, pq.Array(trackID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var trackTabLinks []*models.TrackTabEntry
	for rows.Next() {
		trackTabLink := &models.TrackTabEntry{}
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
