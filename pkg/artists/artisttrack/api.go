package artisttrack

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing artists.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (a *API) LinkArtistToTrack(artistId, trackId string) {
	_, err := a.Database.DB.Exec(insertArtistTrackQuery, artistId, trackId)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistId, trackId, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistId, trackId)
	}
}

// GetArtistToTrackLink retrieves the 'artist to track' link for the provided artist ID.
func (a *API) GetArtistToTrackLink(artistID string) (*ArtistTrack, error) {
	artistTracks, err := a.GetArtistToTrackLinks(artistID)
	if err != nil {
		return nil, err
	}
	return artistTracks[0], err
}

// GetArtistToTrackLinks retrieves the 'artist to track' link for the provided artist IDs.
func (a *API) GetArtistToTrackLinks(artistID ...string) ([]*ArtistTrack, error) {
	rows, err := a.Database.DB.Query(getArtistTrackLinks, artistID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artistTrackLink []*ArtistTrack
	for rows.Next() {
		var artistTrack *ArtistTrack
		err := rows.Scan(&artistTrack.ArtistID, &artistTrack.TrackID)
		if err != nil {
			return nil, err
		}
		artistTrackLink = append(artistTrackLink, artistTrack)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return artistTrackLink, nil
}
