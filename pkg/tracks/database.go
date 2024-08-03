package tracks

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to tracks in the database.
type DataOperations interface {
	InsertTracks(tracks ...*Track)
	InsertTrack(track *Track)
	GetTrack(trackID uuid.UUID) (*Track, error)
	GetTracks(trackID ...uuid.UUID) ([]*Track, error)
}

// DataService is for managing tracks of songs.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertTracks inserts multiple tracks into the tracks table.
func (d *DataService) InsertTracks(tracks ...*Track) {
	for _, track := range tracks {
		d.InsertTrack(track)
	}
}

// InsertTrack inserts a track into the tracks table.
func (d *DataService) InsertTrack(track *Track) {
	_, err := d.Exec(insertTrackQuery, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
}

// GetTrack retrieves the track, without entity references, for the provided ID.
func (d *DataService) GetTrack(trackID uuid.UUID) (*Track, error) {
	tracks, err := d.GetTracks(trackID)
	if err != nil {
		return nil, err
	}
	return tracks[0], nil
}

// GetTracks retrieves the tracks, without entity references, for the provided IDs.
func (d *DataService) GetTracks(trackID ...uuid.UUID) ([]*Track, error) {
	rows, err := d.Query(getTracksFromIDs, trackID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tracks []*Track
	for rows.Next() {
		track := &Track{}
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
