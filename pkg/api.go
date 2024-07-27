package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	arttrk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists/artisttrack"
	diff "github.com/shvdg-dev/tunes-to-tabs-api/pkg/difficulties"
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/endpoints"
	inst "github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	ref "github.com/shvdg-dev/tunes-to-tabs-api/pkg/references"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	tbs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	trktab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks/tracktab"
	usrs "github.com/shvdg-dev/tunes-to-tabs-api/pkg/users"
)

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	*logic.DatabaseManager

	artistsService      art.Operations
	tracksService       trk.Operations
	tabsService         tbs.Operations
	usersService        usrs.Operations
	instrumentsService  inst.Operations
	difficultiesService diff.Operations
	sourcesService      src.Operations
	endpointsService    end.Operations
	referencesService   ref.Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{DatabaseManager: database}
}

// Artists initiates upon first use and returns the artists.Operations.
func (a *API) Artists() art.Operations {
	if a.artistsService == nil {
		artistDatabaseService := art.NewDatabaseService(a.DatabaseManager)
		artistTrackDatabaseService := arttrk.NewDatabaseService(a.DatabaseManager)
		artistTrackService := arttrk.NewService(artistTrackDatabaseService)
		artistMappingService := art.NewMappingService()
		a.artistsService = art.NewService(artistDatabaseService, artistMappingService, artistTrackService, nil)
	}
	return a.artistsService
}

// Tracks instantiates upon first use and returns the tracks.Operations.
func (a *API) Tracks() trk.Operations {
	if a.tracksService == nil {
		trackDatabaseService := trk.NewDatabaseService(a.DatabaseManager)
		trackTabDatabaseService := trktab.NewDatabaseService(a.DatabaseManager)
		trackTabService := trktab.NewService(trackTabDatabaseService)
		trackMappingService := trk.NewMappingService()
		a.tracksService = trk.NewService(trackDatabaseService, trackMappingService, trackTabService, a.Tabs())
	}
	return a.tracksService
}

// Tabs instantiates upon first use and returns the tabs.Operations.
func (a *API) Tabs() tbs.Operations {
	if a.tabsService == nil {
		tabDatabaseService := tbs.NewDatabaseService(a.DatabaseManager)
		tabsMappingService := tbs.NewMappingService()
		tabService := tbs.NewService(tabDatabaseService, tabsMappingService)
		a.tabsService = tabService
	}
	return a.tabsService
}

// Users instantiates upon first use and returns the users.Operations.
func (a *API) Users() usrs.Operations {
	if a.usersService == nil {
		usersDatabaseService := usrs.NewDatabaseService(a.DatabaseManager)
		a.usersService = usrs.NewService(usersDatabaseService)
	}
	return a.usersService
}

// Instruments instantiates upon first use and returns the instruments.Operations.
func (a *API) Instruments() inst.Operations {
	if a.instrumentsService == nil {
		instrumentsDatabaseService := inst.NewDatabaseService(a.DatabaseManager)
		a.instrumentsService = inst.NewService(instrumentsDatabaseService)
	}
	return a.instrumentsService
}

// Difficulties instantiates upon first use and returns the difficulties.Operations.
func (a *API) Difficulties() diff.Operations {
	if a.difficultiesService == nil {
		difficultiesDatabaseService := diff.NewDatabaseService(a.DatabaseManager)
		a.difficultiesService = diff.NewService(difficultiesDatabaseService)
	}
	return a.difficultiesService
}

// Sources instantiates upon first use and returns the sources.Operations.
func (a *API) Sources() src.Operations {
	if a.sourcesService == nil {
		sourceDatabaseService := src.NewDatabaseService(a.DatabaseManager)
		a.sourcesService = src.NewService(sourceDatabaseService)
	}
	return a.sourcesService
}

// Endpoints instantiates upon first use and returns the endpoints.Operations.
func (a *API) Endpoints() end.Operations {
	if a.endpointsService == nil {
		endpointsDatabaseService := end.NewDatabaseService(a.DatabaseManager)
		a.endpointsService = end.NewService(endpointsDatabaseService)
	}
	return a.endpointsService
}

// References instantiates upon first use and returns the references.Operations.
func (a *API) References() ref.Operations {
	if a.referencesService == nil {
		referencesDatabaseService := ref.NewDatabaseService(a.DatabaseManager)
		a.referencesService = ref.NewService(referencesDatabaseService)
	}
	return a.referencesService
}
