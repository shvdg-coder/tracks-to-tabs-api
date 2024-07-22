package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
)

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	*logic.DatabaseManager

	artistsService artists.Operations
	tracksService  tracks.Operations
	tabsService    tabs.Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{DatabaseManager: database}
}

// Artists initiates upon first use and returns the artists.Operations.
func (a *API) Artists() artists.Operations {
	if a.artistsService == nil {
		artistDatabaseService := artists.NewDatabaseService(a.DatabaseManager)
		artistTrackDatabaseService := arttrk.NewDatabaseService(a.DatabaseManager)
		artistTrackService := arttrk.NewService(artistTrackDatabaseService)
		artistMappingService := artists.NewMappingService(artistDatabaseService, artistTrackService, a.Tracks())
		a.artistsService = artists.NewService(artistDatabaseService, artistMappingService, artistTrackService)
	}
	return a.artistsService
}

// Tracks instantiates upon first use and returns the tracks.Operations.
func (a *API) Tracks() tracks.Operations {
	if a.tracksService == nil {
		trackDatabaseService := tracks.NewDatabaseService(a.DatabaseManager)
		trackTabDatabaseService := tracktab.NewDatabaseService(a.DatabaseManager)
		trackTabService := tracktab.NewService(trackTabDatabaseService)
		trackMappingService := tracks.NewMappingService(trackDatabaseService, trackTabService, a.Tabs())
		a.tracksService = tracks.NewService(trackDatabaseService, trackMappingService, trackTabService)
	}
	return a.tracksService
}

// Tabs instantiates upon first use and returns the tabs.Operations.
func (a *API) Tabs() tabs.Operations {
	if a.tabsService == nil {
		tabDatabaseService := tabs.NewDatabaseService(a.DatabaseManager)
		tabService := tabs.NewService(tabDatabaseService)
		a.tabsService = tabService
	}
	return a.tabsService
}
