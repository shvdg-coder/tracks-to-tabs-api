package artists

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing artists.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateArtistsTable creates an artists table if it doesn't already exist.
func (a *API) CreateArtistsTable() {
	_, err := a.Database.DB.Exec(createArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artists' table")
	}
}

// DropArtistsTable drops the artists table if it exists.
func (a *API) DropArtistsTable() {
	_, err := a.Database.DB.Exec(dropArtistsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artists' table")
	}
}

// InsertArtists inserts multiple artists into the artists table.
func (a *API) InsertArtists(artists ...*Artist) {
	for _, artist := range artists {
		a.InsertArtist(artist)
	}
}

// InsertArtist inserts an artist into the artists table.
func (a *API) InsertArtist(artist *Artist) {
	_, err := a.Database.DB.Exec(insertArtistQuery, artist.ID, artist.Name)
	if err != nil {
		log.Printf("Failed inserting user with name '%s': %s", artist.Name, err.Error())
	} else {
		log.Printf("Successfully inserted artist '%s' into the 'artists' table", artist.Name)
	}
}

// GetArtists retrieves artists, without references to other entities.
func (a *API) GetArtists(artistID ...string) ([]*Artist, error) {
	rows, err := a.Database.DB.Query(getArtistsFromIDs, artistID)
	if err != nil {
		return nil, err
	}

	var artists []*Artist

	for rows.Next() {
		artist := &Artist{}
		err := rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}

// GetArtistsCascading retrieves artists, with references to other entities.
func (a *API) GetArtistsCascading(artistID ...string) ([]*Artist, error) {
	//artists, err := a.GetArtists(artistID...)
	//if err != nil {
	//	return nil, err
	//}
	return nil, nil
}
