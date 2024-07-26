package tracks

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to tracks in the database.
type DatabaseOperations interface {
	InsertTracks(tracks ...*Track)
	InsertTrack(track *Track)
	GetTrack(trackID uuid.UUID) (*Track, error)
	GetTracks(trackID ...uuid.UUID) ([]*Track, error)
}

// DatabaseService is for managing tracks of songs.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DatabaseOperations {
	return &DatabaseService{Database: database}
}

// InsertTracks inserts multiple tracks into the tracks table.
func (a *DatabaseService) InsertTracks(tracks ...*Track) {
	for _, track := range tracks {
		a.InsertTrack(track)
	}
}

// InsertTrack inserts a track into the tracks table.
func (a *DatabaseService) InsertTrack(track *Track) {
	_, err := a.Database.DB.Exec(insertTrackQuery, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
}

// GetTrack retrieves the track, without entity references, for the provided ID.
func (a *DatabaseService) GetTrack(trackID uuid.UUID) (*Track, error) {
	tracks, err := a.GetTracks(trackID)
	if err != nil {
		return nil, err
	}
	return tracks[0], nil
}

// GetTracks retrieves the tracks, without entity references, for the provided IDs.
func (a *DatabaseService) GetTracks(trackID ...uuid.UUID) ([]*Track, error) {
	rows, err := a.Database.DB.Query(getTracksFromIDs, trackID)
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
