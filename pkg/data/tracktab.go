package data

import (
	"database/sql"
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TrackTabData represents operations related to 'track to tab' links.
type TrackTabData interface {
	InsertTrackTabEntries(trackTabs ...*models.TrackTabEntry) error
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

// GetTrackToTabEntries retrieves the 'track to tab' links for the provided track IDs.
func (d *TrackTabSvc) GetTrackToTabEntries(trackIDs ...uuid.UUID) ([]*models.TrackTabEntry, error) {
	return logic.BatchGet(d, batchSize, queries.GetTrackTabLinks, trackIDs, scanTrackTabEntry)
}

// scanTrackTabEntry scans a row into a models.TrackTabEntry.
func scanTrackTabEntry(rows *sql.Rows) (*models.TrackTabEntry, error) {
	trackTabLink := &models.TrackTabEntry{}
	if err := rows.Scan(&trackTabLink.TrackID, &trackTabLink.TabID); err != nil {
		return nil, err
	}
	return trackTabLink, nil
}
