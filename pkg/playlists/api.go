package playlists

import logic "github.com/shvdg-dev/base-logic/pkg"

// API is for handling playlists.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}
