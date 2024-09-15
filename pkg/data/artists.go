package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// ArtistData represents operations related to artists in the database.
type ArtistData interface {
	InsertArtistEntries(artist ...*models.ArtistEntry)
	InsertArtistEntry(artist *models.ArtistEntry)
	GetArtistsEntries(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
	GetArtistEntry(artistID uuid.UUID) (*models.ArtistEntry, error)
}

// ArtistSvc is for managing artists.
type ArtistSvc struct {
	logic.DbOperations
}

// NewArtistSvc creates a new instance of the ArtistSvc struct.
func NewArtistSvc(database logic.DbOperations) ArtistData {
	return &ArtistSvc{database}
}

// InsertArtistEntries inserts multiple ArtistEntry's into the artists table.
func (d *ArtistSvc) InsertArtistEntries(artists ...*models.ArtistEntry) {
	for _, artist := range artists {
		d.InsertArtistEntry(artist)
	}
}

// InsertArtistEntry inserts an ArtistEntry into the artists table.
func (d *ArtistSvc) InsertArtistEntry(artist *models.ArtistEntry) {
	_, err := d.Exec(queries.InsertArtist, artist.ID, artist.Name)
	if err != nil {
		log.Printf("Failed inserting user: %s", err.Error())
	}
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
	rows, err := d.Query(queries.GetArtistsFromIDs, pq.Array(artistID))
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
