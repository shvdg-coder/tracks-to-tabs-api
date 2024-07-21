package artisttrack

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseService is for managing artists.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) *DatabaseService {
	return &DatabaseService{Database: database}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (d *DatabaseService) LinkArtistToTrack(artistId, trackId string) {
	_, err := d.Database.DB.Exec(insertArtistTrackQuery, artistId, trackId)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistId, trackId, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistId, trackId)
	}
}

// GetArtistToTrackLink retrieves the 'artist to track' link for the provided artist ID.
func (d *DatabaseService) GetArtistToTrackLink(artistID string) (*ArtistTrack, error) {
	artistTracks, err := d.GetArtistToTrackLinks(artistID)
	if err != nil {
		return nil, err
	}
	return artistTracks[0], err
}

// GetArtistToTrackLinks retrieves the 'artist to track' link for the provided artist IDs.
func (d *DatabaseService) GetArtistToTrackLinks(artistID ...string) ([]*ArtistTrack, error) {
	rows, err := d.Database.DB.Query(getArtistTrackLinks, artistID)
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
