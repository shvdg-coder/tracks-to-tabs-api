package data

import (
	"database/sql"
	"github.com/google/uuid"
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
		data[i] = link.Fields()
	}

	fieldNames := []string{"artist_id", "track_id"}
	return d.BulkInsert("artist_track", fieldNames, data)
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
func (d *ArtistTrackSvc) GetArtistToTrackEntries(artistIDs ...uuid.UUID) ([]*models.ArtistTrackEntry, error) {
	return logic.BatchGet(d, batchSize, queries.GetArtistTrackLinks, artistIDs, scanArtistTrackEntry)
}

// scanArtistTrackEntry scans a row into a models.ArtistTrackEntry.
func scanArtistTrackEntry(rows *sql.Rows) (*models.ArtistTrackEntry, error) {
	artistTrack := &models.ArtistTrackEntry{}
	if err := rows.Scan(&artistTrack.ArtistID, &artistTrack.TrackID); err != nil {
		return nil, err
	}
	return artistTrack, nil
}
