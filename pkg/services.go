package pkg

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/schemas"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/services"
)

// SvcOps represents operations regarding all the services.
type SvcOps interface {
	services.ArtistOps
	services.TrackOps
	services.TabOps
	services.InstrumentOps
	services.DifficultyOps
	services.ReferenceOps
	services.SourceOps
	services.EndpointOps
	services.ArtistTrackOps
	services.TrackTabOps
	services.UserOps
	services.ResourceOps
}

// SvcManager instantiates and handles the different services.
type SvcManager struct {
	services.ArtistOps
	services.TrackOps
	services.TabOps
	services.InstrumentOps
	services.DifficultyOps
	services.ReferenceOps
	services.SourceOps
	services.EndpointOps
	services.ArtistTrackOps
	services.TrackTabOps
	services.UserOps
	services.ResourceOps
}

// NewSvcManager creates a new instance of the SvcManager.
func NewSvcManager(database logic.DbOperations) *SvcManager {
	artistTrackSvc := createArtistTrackService(database)
	trackTabSvc := createTrackTabSvc(database)
	endpointsSvc := createEndpointsSvc(database)
	sourcesSvc := createSourcesSvc(database, endpointsSvc)
	referencesSvc := createReferencesSvc(database, sourcesSvc)
	instrumentsSvc := createInstrumentsSvc(database)
	difficultiesSvc := createDifficultiesSvc(database)
	tabsSvc := createTabsSvc(database, instrumentsSvc, difficultiesSvc, referencesSvc)
	tracksSvc := createTracksSvc(database, trackTabSvc, tabsSvc, referencesSvc)

	return &SvcManager{
		ArtistOps:      createArtistsSvc(database, artistTrackSvc, tracksSvc, referencesSvc),
		TrackOps:       createTracksSvc(database, trackTabSvc, tabsSvc, referencesSvc),
		TabOps:         tabsSvc,
		ArtistTrackOps: artistTrackSvc,
		TrackTabOps:    trackTabSvc,
		UserOps:        createUsersService(database),
		InstrumentOps:  createInstrumentsSvc(database),
		DifficultyOps:  createDifficultiesSvc(database),
		EndpointOps:    endpointsSvc,
		SourceOps:      sourcesSvc,
		ReferenceOps:   referencesSvc,
		ResourceOps:    createResourcesSvc(),
	}
}

// createArtistsSvc creates a services.ArtistOps.
func createArtistsSvc(db logic.DbOperations, artistTrack services.ArtistTrackOps, tracks services.TrackOps, references services.ReferenceOps) services.ArtistOps {
	artistDataSvc := data.NewArtistSvc(db)
	artistSchemaSvc := schemas.NewArtistSvc(db)
	artistMappingSvc := mappers.NewArtistSvc()
	return services.NewArtistSvc(artistSchemaSvc, artistDataSvc, artistMappingSvc, artistTrack, tracks, references)
}

// createTracksSvc creates a services.TrackOps.
func createTracksSvc(db logic.DbOperations, trackTab services.TrackTabOps, tabs services.TabOps, references services.ReferenceOps) services.TrackOps {
	trackSchemaSvc := schemas.NewTrackSvc(db)
	trackDataSvc := data.NewTrackSvc(db)
	trackMappingSvc := mappers.NewTrackSvc()
	return services.NewTrackSvc(trackSchemaSvc, trackDataSvc, trackMappingSvc, trackTab, tabs, references)
}

// createTabsSvc creates a services.TabOps.
func createTabsSvc(db logic.DbOperations, instruments services.InstrumentOps, difficulties services.DifficultyOps, references services.ReferenceOps) services.TabOps {
	tabDataSvc := data.NewTabSvc(db)
	tabSchemaSvc := schemas.NewTabSvc(db)
	tabMappingSvc := mappers.NewTabSvc()
	return services.NewTabSvc(tabSchemaSvc, tabDataSvc, tabMappingSvc, instruments, difficulties, references)
}

// createArtistTrackService creates an services.ArtistTrackOps.
func createArtistTrackService(db logic.DbOperations) services.ArtistTrackOps {
	artistTrackSchemaSvc := schemas.NewArtistTrackSvc(db)
	artistTrackDataSvc := data.NewArtistTrackSvc(db)
	return services.NewArtistTrackSvc(artistTrackSchemaSvc, artistTrackDataSvc)
}

// createTrackTabSvc creates a services.TrackTabOps.
func createTrackTabSvc(db logic.DbOperations) services.TrackTabOps {
	trackTabSchemaSvc := schemas.NewTrackTabSvc(db)
	trackTabDataSvc := data.NewTrackTabSvc(db)
	return services.NewTrackTabSvc(trackTabSchemaSvc, trackTabDataSvc)
}

// createUsersService creates a services.UserOps.
func createUsersService(db logic.DbOperations) services.UserOps {
	userSchemaSvc := schemas.NewUserSvc(db)
	usersDataSvc := data.NewUserSvc(db)
	return services.NewUserSvc(userSchemaSvc, usersDataSvc)
}

// createInstrumentsSvc creates a services.InstrumentOps.
func createInstrumentsSvc(db logic.DbOperations) services.InstrumentOps {
	instrumentSchemaSvc := schemas.NewInstrumentSvc(db)
	instrumentsDataSvc := data.NewInstrumentSvc(db)
	instrumentMappingSvc := mappers.NewInstrumentSvc()
	return services.NewInstrumentSvc(instrumentSchemaSvc, instrumentsDataSvc, instrumentMappingSvc)
}

// createDifficultiesSvc creates a services.DifficultyOps.
func createDifficultiesSvc(db logic.DbOperations) services.DifficultyOps {
	difficultySchemaSvc := schemas.NewDifficultySvc(db)
	difficultiesDataSvc := data.NewDifficultySvc(db)
	difficultiesMappingSvc := mappers.NewDifficultySvc()
	return services.NewDifficultySvc(difficultySchemaSvc, difficultiesDataSvc, difficultiesMappingSvc)
}

// createSourcesSvc creates a services.SourceOps.
func createSourcesSvc(db logic.DbOperations, endpoints services.EndpointOps) services.SourceOps {
	sourceSchemaSvc := schemas.NewSourceSvc(db)
	sourceDataSvc := data.NewSourceSvc(db)
	sourceMappingSvc := mappers.NewSourceSvc()
	return services.NewSourceSvc(sourceSchemaSvc, sourceDataSvc, sourceMappingSvc, endpoints)
}

// createEndpointsSvc creates a services.EndpointOps.
func createEndpointsSvc(db logic.DbOperations) services.EndpointOps {
	endpointSchemaSvc := schemas.NewEndpointSvc(db)
	endpointsDataSvc := data.NewEndpointSvc(db)
	endpointsMappingSvc := mappers.NewEndpointSvc()
	return services.NewEndpointSvc(endpointSchemaSvc, endpointsDataSvc, endpointsMappingSvc)
}

// createReferencesSvc creates a services.ReferenceOps.
func createReferencesSvc(db logic.DbOperations, sources services.SourceOps) services.ReferenceOps {
	referenceSchemaSvc := schemas.NewReferenceSvc(db)
	referencesDataSvc := data.NewReferenceSvc(db)
	referencesMappingSvc := mappers.NewReferenceSvc()
	return services.NewReferenceSvc(referenceSchemaSvc, referencesDataSvc, referencesMappingSvc, sources)
}

// createResourcesSvc creates a services.ResourceSvc.
func createResourcesSvc() services.ResourceOps {
	return services.NewResourceSvc()
}
