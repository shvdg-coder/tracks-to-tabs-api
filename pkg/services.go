package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	art "github.com/shvdg-dev/tunes-to-tabs-api/pkg/services"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

type ArtistsService struct{ art.Operations }
type TracksService struct{ art.Operations }
type TabsService struct{ art.Operations }
type ArtistTrackService struct{ art.Operations }
type TrackTabService struct{ art.Operations }
type UsersService struct{ art.Operations }
type InstrumentsService struct{ art.Operations }
type DifficultiesService struct{ art.Operations }
type SourcesService struct{ art.Operations }
type EndpointsService struct{ art.Operations }
type ReferencesService struct{ art.Operations }

// Operations represents operations regarding all the services.
type Operations interface {
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
	art.Operations
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
	artistDataService := database.NewTabSvc(db)
	artistMappingService := mappers.NewArtistServ()
	return &ArtistsService{art.NewService(artistDataService, artistMappingService, artistTrack, tracks, references)}
}

// createTracksService creates a TracksService.
func createTracksService(db logic.DbOperations, trackTab *TrackTabService, tabs *TabsService, references *ReferencesService) *TracksService {
	trackDataService := trk.NewDataService(db)
	trackMappingService := mappers.NewMappingService()
	return &TracksService{art.NewService(trackDataService, trackMappingService, trackTab, tabs, references)}
}

// createTabsService creates a TabsService.
func createTabsService(db logic.DbOperations, instruments *InstrumentsService, difficulties *DifficultiesService, references *ReferencesService) *TabsService {
	tabDataService := database.NewTabSvc(db)
	tabMappingService := mappers.NewMappingService()
	return &TabsService{art.NewService(tabDataService, tabMappingService, instruments, difficulties, references)}
}

// createArtistTrackService creates an ArtistTrackService.
func createArtistTrackService(db logic.DbOperations) *ArtistTrackService {
	artistTrackDataService := database.NewTabSvc(db)
	return &ArtistTrackService{art.NewService(artistTrackDataService)}
}

// createTrackTabService creates a TrackTabService.
func createTrackTabService(db logic.DbOperations) *TrackTabService {
	trackTabDataService := database.NewTabSvc(db)
	return &TrackTabService{art.NewService(trackTabDataService)}
}

// createUsersService creates a UsersService.
func createUsersService(db logic.DbOperations) *UsersService {
	usersDataService := database.NewUserSvc(db)
	return &UsersService{art.NewService(usersDataService)}
}

// createInstrumentsService creates an InstrumentsService.
func createInstrumentsService(db logic.DbOperations) *InstrumentsService {
	instrumentsDataService := database.NewTabSvc(db)
	return &InstrumentsService{art.NewService(instrumentsDataService)}
}

// createDifficultiesService creates a DifficultiesService.
func createDifficultiesService(db logic.DbOperations) *DifficultiesService {
	difficultiesDataService := database.NewTabSvc(db)
	return &DifficultiesService{art.NewService(difficultiesDataService)}
}

// createSourcesService creates a SourcesService.
func createSourcesService(db logic.DbOperations, endpoints *EndpointsService) *SourcesService {
	sourceDataService := database.NewTabSvc(db)
	sourceMappingService := mappers.NewMappingService()
	return &SourcesService{art.NewService(sourceDataService, sourceMappingService, endpoints)}
}

// createEndpointsService creates an EndpointsService.
func createEndpointsService(db logic.DbOperations) *EndpointsService {
	endpointsDataService := database.NewTabSvc(db)
	return &EndpointsService{art.NewService(endpointsDataService)}
}

// createReferencesService creates a ReferencesService.
func createReferencesService(db logic.DbOperations, sources *SourcesService) *ReferencesService {
	referencesDataService := database.NewTabSvc(db)
	referencesMappingService := mappers.NewMappingService()
	return &ReferencesService{art.NewService(referencesDataService, referencesMappingService, sources)}
}
