package artists

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to artists in the database.
type DatabaseOperations interface {
	InsertArtist(artist *Artist)
	InsertArtists(artist ...*Artist)
	GetArtist(artistID string) (*Artist, error)
	GetArtists(artistID ...string) ([]*Artist, error)
}

// DatabaseService is for managing artists.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DatabaseOperations {
	return &DatabaseService{Database: database}
}

// InsertArtists inserts multiple artists into the artists table.
func (d *DatabaseService) InsertArtists(artists ...*Artist) {
	for _, artist := range artists {
		d.InsertArtist(artist)
	}
}

// InsertArtist inserts an artist into the artists table.
func (d *DatabaseService) InsertArtist(artist *Artist) {
	_, err := d.Database.DB.Exec(insertArtistQuery, artist.ID, artist.Name)
	if err != nil {
		log.Printf("Failed inserting user with name '%s': %s", artist.Name, err.Error())
	} else {
		log.Printf("Successfully inserted artist '%s' into the 'artists' table", artist.Name)
	}
}

// GetArtist retrieves an artist, without entity references, for the provided ID.
func (d *DatabaseService) GetArtist(artistID string) (*Artist, error) {
	artists, err := d.GetArtists(artistID)
	if err != nil {
		return nil, err
	}
	return artists[0], nil
}

// GetArtists retrieves artists, without entity references, for the provided IDs.
func (d *DatabaseService) GetArtists(artistID ...string) ([]*Artist, error) {
	rows, err := d.Database.DB.Query(getArtistsFromIDs, artistID)
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
