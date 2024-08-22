package artisttrack

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to 'artists to tracks' links.
type DataOperations interface {
	LinkArtistToTrack(artistId, trackId uuid.UUID)
	GetArtistToTrackLink(artistID uuid.UUID) (*ArtistTrack, error)
	GetArtistToTrackLinks(artistID ...uuid.UUID) ([]*ArtistTrack, error)
}

// DataService is for managing 'artists to tracks' links.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (d *DataService) LinkArtistToTrack(artistId, trackId uuid.UUID) {
	_, err := d.Exec(insertArtistTrackQuery, artistId, trackId)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistId, trackId, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistId, trackId)
	}
}

// GetArtistToTrackLink retrieves the 'artist to track' link for the provided artist ID.
func (d *DataService) GetArtistToTrackLink(artistID uuid.UUID) (*ArtistTrack, error) {
	artistTracks, err := d.GetArtistToTrackLinks(artistID)
	if err != nil {
		return nil, err
	}
	return artistTracks[0], err
}

// GetArtistToTrackLinks retrieves the 'artist to track' link for the provided artist IDs.
func (d *DataService) GetArtistToTrackLinks(artistID ...uuid.UUID) ([]*ArtistTrack, error) {
	rows, err := d.Query(getArtistTrackLinks, pq.Array(artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artistTrackLink []*ArtistTrack
	for rows.Next() {
		artistTrack := &ArtistTrack{}
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
