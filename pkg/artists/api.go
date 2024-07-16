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

// CreateArtistTrackTable creates an artist_track table if it doesn't already exist.
func (a *API) CreateArtistTrackTable() {
	_, err := a.Database.DB.Exec(createArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artist_track' table")
	}
}

// DropArtistTrackTable drops the artist_track table if it exists.
func (a *API) DropArtistTrackTable() {
	_, err := a.Database.DB.Exec(dropArtistTrackTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artist_track' table")
	}
}

// LinkArtistToTrack inserts a link between an artist and a track into the artist_track table.
func (a *API) LinkArtistToTrack(artistId, trackId string) {
	_, err := a.Database.DB.Exec(insertArtistTrackQuery, artistId, trackId)
	if err != nil {
		log.Printf("Failed linking artist with ID '%s' and track with ID '%s': %s", artistId, trackId, err.Error())
	} else {
		log.Printf("Successfully linked artist with ID '%s' and track with ID '%s'", artistId, trackId)
	}
}

// CreateArtistsTracksTabsView creates a view with the artists, tracks and tabs.
func (a *API) CreateArtistsTracksTabsView() {
	_, err := a.Database.DB.Exec(createArtistsTracksTabsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artists, tracks, tabs' view.")
	}
}

// DropArtistsTracksTabsView drops the view with the artists, tracks and tabs.
func (a *API) DropArtistsTracksTabsView() {
	_, err := a.Database.DB.Exec(dropArtistsTracksTabsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artists, tracks, tabs' view.")
	}
}
