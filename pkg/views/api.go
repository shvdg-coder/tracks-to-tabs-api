package views

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for retrieving data from views.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateArtistsTracksTabsView creates an artists to tabs view.
func (a *API) CreateArtistsTracksTabsView() {
	_, err := a.Database.DB.Exec(createArtistsTracksTabsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'artists' to 'tabs' view.")
	}
}

// DropArtistsTracksToTabsView drops the artists to tabs view if it exists.
func (a *API) DropArtistsTracksToTabsView() {
	_, err := a.Database.DB.Exec(dropArtistsTracksTabsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'artists' to 'tabs' view.")
	}
}

// CreateSourcesToEndpointsView creates the sources to endpoints view.
func (a *API) CreateSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(createSourcesEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' to 'endpoints' view")
	}
}

// DropSourcesToEndpointsView drops the sources to endpoints view if it exists.
func (a *API) DropSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(dropSourcesEndpointViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' to 'endpoints' view")
	}
}
