package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TrackTabData represents operations related to 'track to tab' links.
type TrackTabData interface {
	InsertTrackTabEntries(trackTabs ...*models.TrackTabEntry) error
	GetTrackToTabEntry(trackID uuid.UUID) (*models.TrackTabEntry, error)
	GetTrackToTabEntries(trackID ...uuid.UUID) ([]*models.TrackTabEntry, error)
}

// TrackTabSvc is for managing tracks of songs.
type TrackTabSvc struct {
	logic.DbOps
}

// NewTrackTabSvc creates a new instance of the TrackTabSvc struct.
func NewTrackTabSvc(database logic.DbOps) *TrackTabSvc {
	return &TrackTabSvc{DbOps: database}
}

// InsertTrackTabEntries inserts links between tracks and tabs into the track_tab table.
func (d *TrackTabSvc) InsertTrackTabEntries(trackTabs ...*models.TrackTabEntry) error {
	data := make([][]interface{}, len(trackTabs))

	for i, link := range trackTabs {
		data[i] = link.Fields()
	}

	fieldNames := []string{"track_id", "tab_id"}
	return d.BulkInsert("track_tab", fieldNames, data)
}

// GetTrackToTabEntry retrieves the 'track to tab' link for the provided ID.
func (d *TrackTabSvc) GetTrackToTabEntry(trackID uuid.UUID) (*models.TrackTabEntry, error) {
	trackTabLinks, err := d.GetTrackToTabEntries(trackID)
	if err != nil {
		return nil, err
	}
	return trackTabLinks[0], nil
}

// GetTrackToTabEntries retrieves the 'track to tab' links for the provided track IDs.
func (d *TrackTabSvc) GetTrackToTabEntries(trackID ...uuid.UUID) ([]*models.TrackTabEntry, error) {
	rows, err := d.DB().Query(queries.GetTrackTabLinks, pq.Array(trackID))
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
