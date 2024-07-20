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

// GetArtistsCascading retrieves artists with the provided internal artist IDs, with references to other entities.
func (a *API) GetArtistsCascading(artistID ...string) ([]*Artist, error) {
	artists, err := a.GetArtists(artistID...)
	if err != nil {
		return nil, err
	}
	trackIDs, err := a.ArtistTrack.GetTrackIDs(artistID...)
	if err != nil {
		return nil, err
	}
	_, err = a.Tracks.GetTracksCascading(trackIDs...)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
