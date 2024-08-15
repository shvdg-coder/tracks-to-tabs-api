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

type ArtistsService struct{ art.Operations }
type TracksService struct{ trk.Operations }
type TabsService struct{ tbs.Operations }
type ArtistTrackService struct{ arttrk.Operations }
type TrackTabService struct{ trktab.Operations }
type UsersService struct{ usrs.Operations }
type InstrumentsService struct{ inst.Operations }
type DifficultiesService struct{ diff.Operations }
type SourcesService struct{ src.Operations }
type EndpointsService struct{ end.Operations }
type ReferencesService struct{ ref.Operations }

// Operations represents operations regarding all the services.
type Operations interface {
	art.Operations
	trk.Operations
	tbs.Operations
	arttrk.Operations
	trktab.Operations
	usrs.Operations
	inst.Operations
	diff.Operations
	src.Operations
	end.Operations
	ref.Operations
}

// ServiceManager instantiates and handles the different services.
type ServiceManager struct {
	*ArtistsService
	*TracksService
	*TabsService
	*ArtistTrackService
	*TrackTabService
	*UsersService
	*InstrumentsService
	*DifficultiesService
	*SourcesService
	*EndpointsService
	*ReferencesService
}

// NewServiceManager creates a new instance of the ServiceManager.
func NewServiceManager(database logic.DbOperations) Operations {
	artistTrackService := createArtistTrackService(database)
	trackTabService := createTrackTabService(database)
	tabsService := createTabsService(database)
	tracksService := createTracksService(database, trackTabService, tabsService)

	return &ServiceManager{
		ArtistsService:      createArtistsService(database, artistTrackService, tracksService),
		TracksService:       createTracksService(database, trackTabService, tabsService),
		TabsService:         tabsService,
		ArtistTrackService:  artistTrackService,
		TrackTabService:     trackTabService,
		UsersService:        createUsersService(database),
		InstrumentsService:  createInstrumentsService(database),
		DifficultiesService: createDifficultiesService(database),
		SourcesService:      createSourcesService(database),
		EndpointsService:    createEndpointsService(database),
		ReferencesService:   createReferencesService(database),
	}
}

// createArtistsService creates an ArtistsService.
func createArtistsService(db logic.DbOperations, artistTrack *ArtistTrackService, tracks *TracksService) *ArtistsService {
	artistDataService := art.NewDataService(db)
	artistMappingService := art.NewMappingService()
	return &ArtistsService{art.NewService(artistDataService, artistMappingService, artistTrack, tracks)}
}

// createTracksService creates a TracksService.
func createTracksService(db logic.DbOperations, trackTab *TrackTabService, tabs *TabsService) *TracksService {
	trackDataService := trk.NewDataService(db)
	trackMappingService := trk.NewMappingService()
	return &TracksService{trk.NewService(trackDataService, trackMappingService, trackTab, tabs)}
}

// createTabsService creates a TabsService.
func createTabsService(db logic.DbOperations) *TabsService {
	tabDataService := tbs.NewDataService(db)
	tabMappingService := tbs.NewMappingService()
	return &TabsService{tbs.NewService(tabDataService, tabMappingService)}
}

// createArtistTrackService creates an ArtistTrackService.
func createArtistTrackService(db logic.DbOperations) *ArtistTrackService {
	artistTrackDataService := arttrk.NewDataService(db)
	return &ArtistTrackService{arttrk.NewService(artistTrackDataService)}
}

// createTrackTabService creates a TrackTabService.
func createTrackTabService(db logic.DbOperations) *TrackTabService {
	trackTabDataService := trktab.NewDataService(db)
	return &TrackTabService{trktab.NewService(trackTabDataService)}
}

// createUsersService creates a UsersService.
func createUsersService(db logic.DbOperations) *UsersService {
	usersDataService := usrs.NewDataService(db)
	return &UsersService{usrs.NewService(usersDataService)}
}

// createInstrumentsService creates an InstrumentsService.
func createInstrumentsService(db logic.DbOperations) *InstrumentsService {
	instrumentsDataService := inst.NewDataService(db)
	return &InstrumentsService{inst.NewService(instrumentsDataService)}
}

// createDifficultiesService creates a DifficultiesService.
func createDifficultiesService(db logic.DbOperations) *DifficultiesService {
	difficultiesDataService := diff.NewDataService(db)
	return &DifficultiesService{diff.NewService(difficultiesDataService)}
}

// createSourcesService creates a SourcesService.
func createSourcesService(db logic.DbOperations) *SourcesService {
	sourceDataService := src.NewDataService(db)
	return &SourcesService{src.NewService(sourceDataService)}
}

// createEndpointsService creates an EndpointsService.
func createEndpointsService(db logic.DbOperations) *EndpointsService {
	endpointsDataService := end.NewDataService(db)
	return &EndpointsService{end.NewService(endpointsDataService)}
}

// createReferencesService creates a ReferencesService.
func createReferencesService(db logic.DbOperations) *ReferencesService {
	referencesDataService := ref.NewDataService(db)
	return &ReferencesService{ref.NewService(referencesDataService)}
}
