package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// ArtistTrackData represents operations related to 'artists to tracks' links.
type ArtistTrackData interface {
	InsertArtistTrackEntries(artistTracks ...*models.ArtistTrackEntry) error
	GetArtistToTrackEntry(ID uuid.UUID) (*models.ArtistTrackEntry, error)
	GetArtistToTrackEntries(IDs ...uuid.UUID) ([]*models.ArtistTrackEntry, error)
}

// ArtistTrackSvc is for managing 'artists to tracks' links.
type ArtistTrackSvc struct {
	logic.DbOps
}

// NewArtistTrackSvc creates a new instance of the ArtistTrackSvc struct.
func NewArtistTrackSvc(database logic.DbOps) ArtistTrackData {
	return &ArtistTrackSvc{DbOps: database}
}

// InsertArtistTrackEntries inserts links between artists and tracks into the artist_track table.
func (d *ArtistTrackSvc) InsertArtistTrackEntries(artistTracks ...*models.ArtistTrackEntry) error {
	data := make([][]interface{}, len(artistTracks))

	for i, link := range artistTracks {
		data[i] = logic.GetFields(link)
	}

	return d.BulkInsert("artist_track", logic.GetFieldNames("db", &models.ArtistTrackEntry{}), data)
}

// GetArtistToTrackEntry retrieves the 'artist to track' link for the provided artist or track IDs.
func (d *ArtistTrackSvc) GetArtistToTrackEntry(artistID uuid.UUID) (*models.ArtistTrackEntry, error) {
	artistTracks, err := d.GetArtistToTrackEntries(artistID)
	if err != nil {
		return nil, err
	}
	return artistTracks[0], err
}

// GetArtistToTrackEntries retrieves the 'artist to track' link for the provided artist or track IDs.
func (d *ArtistTrackSvc) GetArtistToTrackEntries(artistID ...uuid.UUID) ([]*models.ArtistTrackEntry, error) {
	rows, err := d.DB().Query(queries.GetArtistTrackLinks, pq.Array(artistID))
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
