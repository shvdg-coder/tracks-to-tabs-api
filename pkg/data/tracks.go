package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// TrackData represents operations related to tracks in the database.
type TrackData interface {
	InsertTracks(tracks ...*models.TrackEntry)
	InsertTrack(track *models.TrackEntry)
	GetTrack(trackID uuid.UUID) (*models.TrackEntry, error)
	GetTracks(trackID ...uuid.UUID) ([]*models.TrackEntry, error)
}

// TrackSvc is for managing tracks of songs.
type TrackSvc struct {
	logic.DbOperations
}

// NewTrackSvc creates a new instance of the TrackSvc struct.
func NewTrackSvc(database logic.DbOperations) TrackData {
	return &TrackSvc{DbOperations: database}
}

// InsertTracks inserts multiple tracks into the tracks table.
func (d *TrackSvc) InsertTracks(tracks ...*models.TrackEntry) {
	for _, track := range tracks {
		d.InsertTrack(track)
	}
}

// InsertTrack inserts a track into the tracks table.
func (d *TrackSvc) InsertTrack(track *models.TrackEntry) {
	_, err := d.Exec(queries.InsertTrack, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
}

// GetTrack retrieves the track, without entity references, for the provided ID.
func (d *TrackSvc) GetTrack(trackID uuid.UUID) (*models.TrackEntry, error) {
	tracks, err := d.GetTracks(trackID)
	if err != nil {
		return nil, err
	}
	return tracks[0], nil
}

// GetTracks retrieves the tracks, without entity references, for the provided IDs.
func (d *TrackSvc) GetTracks(trackID ...uuid.UUID) ([]*models.TrackEntry, error) {
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
