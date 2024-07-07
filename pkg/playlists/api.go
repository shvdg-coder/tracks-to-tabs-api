package playlists

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
)

// API is for handling playlists.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// GetArtists retrieves the artists, present in the playlist, from the database.
func (a *API) GetArtists(playlistURL string) []*art.Artist {
	return nil
}
