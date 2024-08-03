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
	logic.DbOperations

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
func NewAPI(database logic.DbOperations) *API {
	return &API{DbOperations: database}
}

// Artists initiates upon first use and returns the artists.Operations.
func (a *API) Artists() art.Operations {
	if a.artistsService == nil {
		artistMappingService := art.NewMappingService()
		artistDataService := art.NewDataService(a.DbOperations)
		artistTrackDataService := arttrk.NewDataService(a.DbOperations)
		artistTrackService := arttrk.NewService(artistTrackDataService)
		a.artistsService = art.NewService(artistDataService, artistMappingService, artistTrackService, a.Tracks())
	}
	return a.artistsService
}

// Tracks instantiates upon first use and returns the tracks.Operations.
func (a *API) Tracks() trk.Operations {
	if a.tracksService == nil {
		trackMappingService := trk.NewMappingService()
		trackDataService := trk.NewDataService(a.DbOperations)
		trackTabDataService := trktab.NewDataService(a.DbOperations)
		trackTabService := trktab.NewService(trackTabDataService)
		a.tracksService = trk.NewService(trackDataService, trackMappingService, trackTabService, a.Tabs())
	}
	return a.tracksService
}

// Tabs instantiates upon first use and returns the tabs.Operations.
func (a *API) Tabs() tbs.Operations {
	if a.tabsService == nil {
		tabMappingService := tbs.NewMappingService()
		tabDataService := tbs.NewDataService(a.DbOperations)
		a.tabsService = tbs.NewService(tabDataService, tabMappingService)
	}
	return a.tabsService
}

// Users instantiates upon first use and returns the users.Operations.
func (a *API) Users() usrs.Operations {
	if a.usersService == nil {
		usersDataService := usrs.NewDataService(a.DbOperations)
		a.usersService = usrs.NewService(usersDataService)
	}
	return a.usersService
}

// Instruments instantiates upon first use and returns the instruments.Operations.
func (a *API) Instruments() inst.Operations {
	if a.instrumentsService == nil {
		instrumentsDataService := inst.NewDataService(a.DbOperations)
		a.instrumentsService = inst.NewService(instrumentsDataService)
	}
	return a.instrumentsService
}

// Difficulties instantiates upon first use and returns the difficulties.Operations.
func (a *API) Difficulties() diff.Operations {
	if a.difficultiesService == nil {
		difficultiesDataService := diff.NewDataService(a.DbOperations)
		a.difficultiesService = diff.NewService(difficultiesDataService)
	}
	return a.difficultiesService
}

// Sources instantiates upon first use and returns the sources.Operations.
func (a *API) Sources() src.Operations {
	if a.sourcesService == nil {
		sourceDataService := src.NewDataService(a.DbOperations)
		a.sourcesService = src.NewService(sourceDataService)
	}
	return a.sourcesService
}

// Endpoints instantiates upon first use and returns the endpoints.Operations.
func (a *API) Endpoints() end.Operations {
	if a.endpointsService == nil {
		endpointsDataService := end.NewDataService(a.DbOperations)
		a.endpointsService = end.NewService(endpointsDataService)
	}
	return a.endpointsService
}

// References instantiates upon first use and returns the references.Operations.
func (a *API) References() ref.Operations {
	if a.referencesService == nil {
		referencesDataService := ref.NewDataService(a.DbOperations)
		a.referencesService = ref.NewService(referencesDataService)
	}
	return a.referencesService
}
