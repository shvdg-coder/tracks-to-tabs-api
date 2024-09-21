package data

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
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
	txn, err := d.DbOps.DB().Begin()
	if err != nil {
		return fmt.Errorf("failed starting transaction: %w", err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("artists", "id", "name"))
	if err != nil {
		return fmt.Errorf("failed preparing statement: %w", err)
	}

	for _, artist := range artists {
		_, err := stmt.Exec(artist.ID, artist.Name)
		if err != nil {
			return fmt.Errorf("failed inserting artist: %w", err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("failed executing statement: %w", err)
	}

	err = stmt.Close()
	if err != nil {
		return fmt.Errorf("failed closing statement: %w", err)
	}

	err = txn.Commit()
	if err != nil {
		return fmt.Errorf("failed committing transaction: %w", err)
	}

	return nil
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
func (d *ArtistSvc) GetArtistsEntries(artistID ...uuid.UUID) ([]*models.ArtistEntry, error) {
	rows, err := d.DB().Query(queries.GetArtistsFromIDs, pq.Array(artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artists []*models.ArtistEntry
	for rows.Next() {
		artistEntry := &models.ArtistEntry{}
		err := rows.Scan(&artistEntry.ID, &artistEntry.Name)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artistEntry)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return artists, nil
}
