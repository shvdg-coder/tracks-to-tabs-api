package artists

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	at "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	trcks "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"log"
)

// API is for managing artists.
type API struct {
	Database    *logic.DatabaseManager
	ArtistTrack *at.API
	Tracks      *trcks.API
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager, artistTrack *at.API, tracks *trcks.API) *API {
	return &API{Database: database, ArtistTrack: artistTrack, Tracks: tracks}
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

// GetArtist retrieves an artist, without entity references, for the provided ID.
func (a *API) GetArtist(artistID string) (*Artist, error) {
	artists, err := a.GetArtists(artistID)
	if err != nil {
		return nil, err
	}
	return artists[0], nil
}

// GetArtists retrieves artists, without entity references, for the provided IDs.
func (a *API) GetArtists(artistID ...string) ([]*Artist, error) {
	rows, err := a.Database.DB.Query(getArtistsFromIDs, artistID)
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
