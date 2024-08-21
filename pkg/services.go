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
	endpointsService := createEndpointsService(database)
	sourcesService := createSourcesService(database, endpointsService)
	referencesService := createReferencesService(database, sourcesService)
	instrumentsService := createInstrumentsService(database)
	difficultiesService := createDifficultiesService(database)
	tabsService := createTabsService(database, instrumentsService, difficultiesService, referencesService)
	tracksService := createTracksService(database, trackTabService, tabsService, referencesService)

	return &ServiceManager{
		ArtistsService:      createArtistsService(database, artistTrackService, tracksService, referencesService),
		TracksService:       createTracksService(database, trackTabService, tabsService, referencesService),
		TabsService:         tabsService,
		ArtistTrackService:  artistTrackService,
		TrackTabService:     trackTabService,
		UsersService:        createUsersService(database),
		InstrumentsService:  createInstrumentsService(database),
		DifficultiesService: createDifficultiesService(database),
		EndpointsService:    endpointsService,
		SourcesService:      sourcesService,
		ReferencesService:   referencesService,
	}
}

// createArtistsService creates an ArtistsService.
func createArtistsService(db logic.DbOperations, artistTrack *ArtistTrackService, tracks *TracksService, references *ReferencesService) *ArtistsService {
	artistDataService := art.NewDataService(db)
	artistMappingService := art.NewMappingService()
	return &ArtistsService{art.NewService(artistDataService, artistMappingService, artistTrack, tracks, references)}
}

// createTracksService creates a TracksService.
func createTracksService(db logic.DbOperations, trackTab *TrackTabService, tabs *TabsService, references *ReferencesService) *TracksService {
	trackDataService := trk.NewDataService(db)
	trackMappingService := trk.NewMappingService()
	return &TracksService{trk.NewService(trackDataService, trackMappingService, trackTab, tabs, references)}
}

// createTabsService creates a TabsService.
func createTabsService(db logic.DbOperations, instruments *InstrumentsService, difficulties *DifficultiesService, references *ReferencesService) *TabsService {
	tabDataService := tbs.NewDataService(db)
	tabMappingService := tbs.NewMappingService()
	return &TabsService{tbs.NewService(tabDataService, tabMappingService, instruments, difficulties, references)}
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
func createSourcesService(db logic.DbOperations, endpoints *EndpointsService) *SourcesService {
	sourceDataService := src.NewDataService(db)
	sourceMappingService := src.NewMappingService()
	return &SourcesService{src.NewService(sourceDataService, sourceMappingService, endpoints)}
}

// createEndpointsService creates an EndpointsService.
func createEndpointsService(db logic.DbOperations) *EndpointsService {
	endpointsDataService := end.NewDataService(db)
	return &EndpointsService{end.NewService(endpointsDataService)}
}

// createReferencesService creates a ReferencesService.
func createReferencesService(db logic.DbOperations, sources *SourcesService) *ReferencesService {
	referencesDataService := ref.NewDataService(db)
	referencesMappingService := ref.NewMappingService()
	return &ReferencesService{ref.NewService(referencesDataService, referencesMappingService, sources)}
}
