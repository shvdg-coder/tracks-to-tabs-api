package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
)

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	*logic.DatabaseManager

	artistsService artists.Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{DatabaseManager: database}
}

// Artists initiates upon first use and returns the artists.Operations.
func (a *API) Artists() artists.Operations {
	if a.artistsService == nil {
		artistDatabaseService := artists.NewDatabaseService(a.DatabaseManager)
		artistMappingService := artists.NewMappingService(artistDatabaseService, nil, nil)
		artistTrackDatabaseService := arttrk.NewDatabaseService(a.DatabaseManager)
		artistTrackService := arttrk.NewService(artistTrackDatabaseService)
		a.artistsService = artists.NewService(artistDatabaseService, artistMappingService, artistTrackService)
	}
	return a.artistsService
}
