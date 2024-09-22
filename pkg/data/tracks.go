package data

import (
	"database/sql"
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TrackData represents operations related to tracks in the database.
type TrackData interface {
	InsertTrackEntries(tracks ...*models.TrackEntry) error
	GetTrackEntries(trackID ...uuid.UUID) ([]*models.TrackEntry, error)
}

// TrackSvc is for managing tracks of songs.
type TrackSvc struct {
	logic.DbOps
}

// NewTrackSvc creates a new instance of the TrackSvc struct.
func NewTrackSvc(database logic.DbOps) TrackData {
	return &TrackSvc{DbOps: database}
}

// InsertTrackEntries inserts multiple tracks into the tracks table.
func (d *TrackSvc) InsertTrackEntries(tracks ...*models.TrackEntry) error {
	data := make([][]interface{}, len(tracks))

	for i, track := range tracks {
		data[i] = track.Fields()
	}

	fieldNames := []string{"id", "title", "duration"}
	return d.BulkInsert("tracks", fieldNames, data)
}

// GetTrackEntries retrieves track entries, without entity references, for the provided IDs.
func (d *TrackSvc) GetTrackEntries(trackIDs ...uuid.UUID) ([]*models.TrackEntry, error) {
	return logic.BatchGet(d, batchSize, queries.GetTracksFromIDs, trackIDs, scanTrackEntry)
}

// scanTrackEntry scans a row into a models.TrackEntry.
func scanTrackEntry(rows *sql.Rows) (*models.TrackEntry, error) {
	track := &models.TrackEntry{}
	if err := rows.Scan(&track.ID, &track.Title, &track.Duration); err != nil {
		return nil, err
	}
	return track, nil
}
