package views

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing 'artists, tracks, tabs' views.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
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
