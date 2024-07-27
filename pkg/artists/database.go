package artists

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to artists in the database.
type DataOperations interface {
	InsertArtist(artist *Artist)
	InsertArtists(artist ...*Artist)
	GetArtist(artistID uuid.UUID) (*Artist, error)
	GetArtists(artistID ...uuid.UUID) ([]*Artist, error)
}

// DataService is for managing artists.
type DataService struct {
	*logic.DatabaseManager
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database *logic.DatabaseManager) DataOperations {
	return &DataService{database}
}

// InsertArtists inserts multiple artists into the artists table.
func (d *DataService) InsertArtists(artists ...*Artist) {
	for _, artist := range artists {
		d.InsertArtist(artist)
	}
}

// InsertArtist inserts an artist into the artists table.
func (d *DataService) InsertArtist(artist *Artist) {
	_, err := d.DB.Exec(insertArtistQuery, artist.ID, artist.Name)
	if err != nil {
		log.Printf("Failed inserting user with name '%s': %s", artist.Name, err.Error())
	} else {
		log.Printf("Successfully inserted artist '%s' into the 'artists' table", artist.Name)
	}
}

// GetArtist retrieves an artist, without entity references, for the provided ID.
func (d *DataService) GetArtist(artistID uuid.UUID) (*Artist, error) {
	artists, err := d.GetArtists(artistID)
	if err != nil {
		return nil, err
	}
	return artists[0], nil
}

// GetArtists retrieves artists, without entity references, for the provided IDs.
func (d *DataService) GetArtists(artistID ...uuid.UUID) ([]*Artist, error) {
	rows, err := d.DB.Query(getArtistsFromIDs, artistID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var artists []*Artist
	for rows.Next() {
		artist := &Artist{}
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
