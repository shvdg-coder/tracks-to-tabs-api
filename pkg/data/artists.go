package data

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// ArtistData represents operations related to artists in the database.
type ArtistData interface {
	InsertArtistEntries(artist ...*models.ArtistEntry) error
	GetArtistsEntries(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
	GetArtistEntry(artistID uuid.UUID) (*models.ArtistEntry, error)
}

// ArtistSvc is for managing artists.
type ArtistSvc struct {
	logic.DbOps
}

// NewArtistSvc creates a new instance of the ArtistSvc struct.
func NewArtistSvc(database logic.DbOps) ArtistData {
	return &ArtistSvc{database}
}

// InsertArtistEntries inserts multiple ArtistEntry's into the artists table using bulk insert.
func (d *ArtistSvc) InsertArtistEntries(artists ...*models.ArtistEntry) error {
	data := make([][]interface{}, len(artists))

	for i, artist := range artists {
		data[i] = artist.Fields()
	}

	fieldNames := []string{"id", "name"}
	return d.BulkInsert("artists", fieldNames, data)
}

// GetArtistEntry retrieves an artist entry, without entity references, for the provided ID.
func (d *ArtistSvc) GetArtistEntry(artistID uuid.UUID) (*models.ArtistEntry, error) {
	artists, err := d.GetArtistsEntries(artistID)
	if err != nil {
		return nil, err
	}
	return artists[0], nil
}

// GetArtistsEntries retrieves artist entries, without entity references, for the provided IDs.
func (d *ArtistSvc) GetArtistsEntries(artistIDs ...uuid.UUID) ([]*models.ArtistEntry, error) {
	return logic.BatchGet(d, 1000, queries.GetArtistsFromIDs, artistIDs, scanArtistEntry)
}

// scanArtistEntry scans a row into a models.ArtistEntry.
func scanArtistEntry(rows *sql.Rows) (*models.ArtistEntry, error) {
	artistEntry := &models.ArtistEntry{}
	if err := rows.Scan(&artistEntry.ID, &artistEntry.Name); err != nil {
		return nil, err
	}
	return artistEntry, nil
}
