package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// TrackData represents operations related to tracks in the database.
type TrackData interface {
	InsertTrackEntry(track *models.TrackEntry)
	InsertTrackEntries(tracks ...*models.TrackEntry)
	GetTrackEntry(trackID uuid.UUID) (*models.TrackEntry, error)
	GetTrackEntries(trackID ...uuid.UUID) ([]*models.TrackEntry, error)
}

// TrackSvc is for managing tracks of songs.
type TrackSvc struct {
	logic.DbOperations
}

// NewTrackSvc creates a new instance of the TrackSvc struct.
func NewTrackSvc(database logic.DbOperations) TrackData {
	return &TrackSvc{DbOperations: database}
}

// InsertTrackEntries inserts multiple tracks into the tracks table.
func (d *TrackSvc) InsertTrackEntries(tracks ...*models.TrackEntry) {
	for _, track := range tracks {
		d.InsertTrackEntry(track)
	}
}

// InsertTrackEntry inserts a track into the tracks table.
func (d *TrackSvc) InsertTrackEntry(track *models.TrackEntry) {
	_, err := d.Exec(queries.InsertTrack, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track: %s", err.Error())
	}
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
	rows, err := d.Query(queries.GetTracksFromIDs, pq.Array(trackID))
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
