package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// TrackData represents operations related to tracks in the database.
type TrackData interface {
	InsertTrackEntries(tracks ...*models.TrackEntry) error
	GetTrackEntry(trackID uuid.UUID) (*models.TrackEntry, error)
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

// GetTrackEntry retrieves a track entry, without entity references, for the provided ID.
func (d *TrackSvc) GetTrackEntry(trackID uuid.UUID) (*models.TrackEntry, error) {
	tracks, err := d.GetTrackEntries(trackID)
	if err != nil {
		return nil, err
	}
	return tracks[0], nil
}

// GetTrackEntries retrieves tracks entries, without entity references, for the provided IDs.
func (d *TrackSvc) GetTrackEntries(trackID ...uuid.UUID) ([]*models.TrackEntry, error) {
	rows, err := d.DB().Query(queries.GetTracksFromIDs, pq.Array(trackID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tracks []*models.TrackEntry
	for rows.Next() {
		track := &models.TrackEntry{}
		err := rows.Scan(&track.ID, &track.Title, &track.Duration)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tracks, nil
}
