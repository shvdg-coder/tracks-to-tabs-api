package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// ArtistOps represents operations related to artists in the database.
type ArtistOps interface {
	InsertArtist(artist *models.ArtistEntry)
	InsertArtists(artist ...*models.ArtistEntry)
	GetArtist(artistID uuid.UUID) (*models.ArtistEntry, error)
	GetArtists(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
}

// ArtistServ is for managing artists.
type ArtistServ struct {
	logic.DbOperations
}

// NewArtistServ creates a new instance of the ArtistServ struct.
func NewArtistServ(database logic.DbOperations) ArtistOps {
	return &ArtistServ{database}
}

// InsertArtists inserts multiple ArtistEntry's into the artists table.
func (d *ArtistServ) InsertArtists(artists ...*models.ArtistEntry) {
	for _, artist := range artists {
		d.InsertArtist(artist)
	}
}

// InsertArtist inserts an ArtistEntry into the artists table.
func (d *ArtistServ) InsertArtist(artist *models.ArtistEntry) {
	_, err := d.Exec(queries.InsertArtist, artist.ID, artist.Name)
	if err != nil {
		log.Printf("Failed inserting user with name '%s': %s", artist.Name, err.Error())
	} else {
		log.Printf("Successfully inserted artist '%s' into the 'artists' table", artist.Name)
	}
}

// GetArtist retrieves an artist entry, without entity references, for the provided ID.
func (d *ArtistServ) GetArtist(artistID uuid.UUID) (*models.ArtistEntry, error) {
	artists, err := d.GetArtists(artistID)
	if err != nil {
		return nil, err
	}
	return artists[0], nil
}

// GetArtists retrieves artist entries, without entity references, for the provided IDs.
func (d *ArtistServ) GetArtists(artistID ...uuid.UUID) ([]*models.ArtistEntry, error) {
	rows, err := d.Query(queries.GetArtistsFromIDs, pq.Array(artistID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artists []*models.ArtistEntry
	for rows.Next() {
		artist := &models.ArtistEntry{}
		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return artists, nil
}
