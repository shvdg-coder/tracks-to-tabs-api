package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// ArtistTrackData represents operations related to 'artists to tracks' links.
type ArtistTrackData interface {
	LinkArtistToTrack(artistId, trackId uuid.UUID)
	GetArtistToTrackLink(artistID uuid.UUID) (*models.ArtistTrackEntry, error)
	GetArtistToTrackLinks(artistID ...uuid.UUID) ([]*models.ArtistTrackEntry, error)
}

// ArtistTrackSvc is for managing 'artists to tracks' links.
type ArtistTrackSvc struct {
	logic.DbOperations
}

// NewArtistTrackSvc creates a new instance of the ArtistTrackSvc struct.
func NewArtistTrackSvc(database logic.DbOperations) ArtistTrackData {
	return &ArtistTrackSvc{DbOperations: database}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (d *ArtistTrackSvc) LinkArtistToTrack(artistID, trackID uuid.UUID) {
	_, err := d.Exec(queries.InsertArtistTrack, artistID, trackID)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistID, trackID, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistID, trackID)
	}
}

// GetArtistToTrackLink retrieves the 'artist to track' link for the provided artist ID.
func (d *ArtistTrackSvc) GetArtistToTrackLink(artistID uuid.UUID) (*models.ArtistTrackEntry, error) {
	artistTracks, err := d.GetArtistToTrackLinks(artistID)
	if err != nil {
		return nil, err
	}
	return artistTracks[0], err
}

// GetArtistToTrackLinks retrieves the 'artist to track' link for the provided artist IDs.
func (d *ArtistTrackSvc) GetArtistToTrackLinks(artistID ...uuid.UUID) ([]*models.ArtistTrackEntry, error) {
	rows, err := d.Query(queries.GetArtistTrackLinks, pq.Array(artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artistTrackLink []*models.ArtistTrackEntry
	for rows.Next() {
		artistTrack := &models.ArtistTrackEntry{}
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
