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

// DataOperations represents all API data operations.
type DataOperations interface {
	art.DataOperations
	trk.DataOperations
	tbs.DataOperations
	arttrk.DataOperations
	trktab.DataOperations
	usrs.DataOperations
	inst.DataOperations
	diff.DataOperations
	src.DataOperations
	end.DataOperations
	ref.DataOperations
}

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	DataOperations
	logic.DbOperations

	artistsService      art.Operations
	tracksService       trk.Operations
	tabsService         tbs.Operations
	artistTrackService  arttrk.Operations
	trackTabService     trktab.Operations
	usersService        usrs.Operations
	instrumentsService  inst.Operations
	difficultiesService diff.Operations
	sourcesService      src.Operations
	endpointsService    end.Operations
	referencesService   ref.Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database logic.DbOperations) DataOperations {
	api := &API{DbOperations: database}
	api.initServices()
	return api
}

// initServices initializes the services for the API.
func (a *API) initServices() {
	a.initArtistTrack()
	a.initTrackTab()
	a.initTabs()
	a.initTracks()
	a.initArtists()
	a.initInstruments()
	a.initDifficulties()
	a.initSources()
	a.initEndpoints()
	a.initReferences()
	a.initUsers()
}

// initArtists initiates an artists.Service during the creation of the API.
func (a *API) initArtists() {
	artistDataService := art.NewDataService(a.DbOperations)
	artistMappingService := art.NewMappingService()
	a.artistsService = art.NewService(artistDataService, artistMappingService, a.artistTrackService, a.tracksService)
}

// initTracks initializes the track service during the creation of the API.
func (a *API) initTracks() {
	trackDataService := trk.NewDataService(a.DbOperations)
	trackMappingService := trk.NewMappingService()
	a.tracksService = trk.NewService(trackDataService, trackMappingService, a.trackTabService, a.tabsService)
}

// initTabs initializes the tab service at the creation of the API.
func (a *API) initTabs() {
	tabDataService := tbs.NewDataService(a.DbOperations)
	tabMappingService := tbs.NewMappingService()
	a.tabsService = tbs.NewService(tabDataService, tabMappingService)
}

// initArtistTrack initializes the 'artist to track' link service at the creation of the API.
func (a *API) initArtistTrack() {
	artistTrackDataService := arttrk.NewDataService(a.DbOperations)
	a.artistTrackService = arttrk.NewService(artistTrackDataService)
}

// initTrackTab initializes the 'track to tab' link service at the creation of the API.
func (a *API) initTrackTab() {
	trackTabDataService := trktab.NewDataService(a.DbOperations)
	a.trackTabService = trktab.NewService(trackTabDataService)
}

// initUsers initializes the user service during the creation of the API.
func (a *API) initUsers() {
	usersDataService := usrs.NewDataService(a.DbOperations)
	a.usersService = usrs.NewService(usersDataService)
}

// initInstruments initializes the instrument service at the creation of the API.
func (a *API) initInstruments() {
	instrumentsDataService := inst.NewDataService(a.DbOperations)
	a.instrumentsService = inst.NewService(instrumentsDataService)
}

// initDifficulties initializes the difficulty service during the creation of the API.
func (a *API) initDifficulties() {
	difficultiesDataService := diff.NewDataService(a.DbOperations)
	a.difficultiesService = diff.NewService(difficultiesDataService)
}

// initSources initializes the source service during the creation of the API.
func (a *API) initSources() {
	sourceDataService := src.NewDataService(a.DbOperations)
	a.sourcesService = src.NewService(sourceDataService)
}

// initEndpoints initializes the endpoint service at the creation of the API.
func (a *API) initEndpoints() {
	endpointsDataService := end.NewDataService(a.DbOperations)
	a.endpointsService = end.NewService(endpointsDataService)
}

// initReferences initializes the reference service during creation of the API.
func (a *API) initReferences() {
	referencesDataService := ref.NewDataService(a.DbOperations)
	a.referencesService = ref.NewService(referencesDataService)
}
