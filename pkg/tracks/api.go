package tracks

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trcktb "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
	"log"
)

// API is for managing tracks of songs.
type API struct {
	Database    *logic.DatabaseManager
	TrackTabAPI *trcktb.API
	TabsAPI     *tbs.API
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager, trackTabAPI *trcktb.API, tabsAPI *tbs.API) *API {
	return &API{Database: database, TrackTabAPI: trackTabAPI, TabsAPI: tabsAPI}
}

// InsertTracks inserts multiple tracks into the tracks table.
func (a *API) InsertTracks(tracks ...*Track) {
	for _, track := range tracks {
		a.InsertTrack(track)
	}
}

// InsertTrack inserts a track into the tracks table.
func (a *API) InsertTrack(track *Track) {
	_, err := a.Database.DB.Exec(insertTrackQuery, track.ID, track.Title, track.Duration)
	if err != nil {
		log.Printf("Failed to insert track with title '%s': %s", track.Title, err.Error())
	} else {
		log.Printf("Successfully inserted track into the 'tracks' table with title '%s'", track.Title)
	}
}

// GetTrack retrieves the track, without entity references, for the provided ID.
func (a *API) GetTrack(trackID string) (*Track, error) {
	tracks, err := a.GetTracks(trackID)
	if err != nil {
		return nil, err
	}
	return tracks[0], nil
}

// GetTracks retrieves the tracks, without entity references, for the provided IDs.
func (a *API) GetTracks(trackID ...string) ([]*Track, error) {
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
