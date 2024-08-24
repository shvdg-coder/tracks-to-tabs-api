package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/services"
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
}

// NewSvcManager creates a new instance of the SvcManager.
func NewSvcManager(database logic.DbOperations) SvcOps {
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
	}
}

// createArtistsSvc creates a services.ArtistOps.
func createArtistsSvc(db logic.DbOperations, artistTrack services.ArtistTrackOps, tracks services.TrackOps, references services.ReferenceOps) services.ArtistOps {
	artistDataSvc := data.NewArtistSvc(db)
	artistMappingSvc := mappers.NewArtistSvc()
	return services.NewArtistSvc(artistDataSvc, artistMappingSvc, artistTrack, tracks, references)
}

// createTracksSvc creates a services.TrackOps.
func createTracksSvc(db logic.DbOperations, trackTab services.TrackTabOps, tabs services.TabOps, references services.ReferenceOps) services.TrackOps {
	trackDataSvc := data.NewTrackSvc(db)
	trackMappingSvc := mappers.NewTrackSvc()
	return services.NewTrackSvc(trackDataSvc, trackMappingSvc, trackTab, tabs, references)
}

// createTabsSvc creates a services.TabOps.
func createTabsSvc(db logic.DbOperations, instruments services.InstrumentOps, difficulties services.DifficultyOps, references services.ReferenceOps) services.TabOps {
	tabDataSvc := data.NewTabSvc(db)
	tabMappingSvc := mappers.NewTabSvc()
	return services.NewTabSvc(tabDataSvc, tabMappingSvc, instruments, difficulties, references)
}

// createArtistTrackService creates an services.ArtistTrackOps.
func createArtistTrackService(db logic.DbOperations) services.ArtistTrackOps {
	artistTrackDataSvc := data.NewArtistTrackSvc(db)
	return services.NewArtistTrackSvc(artistTrackDataSvc)
}

// createTrackTabSvc creates a services.TrackTabOps.
func createTrackTabSvc(db logic.DbOperations) services.TrackTabOps {
	trackTabDataSvc := data.NewTrackTabSvc(db)
	return services.NewTrackTabSvc(trackTabDataSvc)
}

// createUsersService creates a services.UserOps.
func createUsersService(db logic.DbOperations) services.UserOps {
	usersDataSvc := data.NewUserSvc(db)
	return services.NewUserSvc(usersDataSvc)
}

// createInstrumentsSvc creates a services.InstrumentOps.
func createInstrumentsSvc(db logic.DbOperations) services.InstrumentOps {
	instrumentsDataSvc := data.NewInstrumentSvc(db)
	instrumentMappingSvc := mappers.NewInstrumentSvc()
	return services.NewInstrumentSvc(instrumentsDataSvc, instrumentMappingSvc)
}

// createDifficultiesSvc creates a services.DifficultyOps.
func createDifficultiesSvc(db logic.DbOperations) services.DifficultyOps {
	difficultiesDataSvc := data.NewDifficultySvc(db)
	difficultiesMappingSvc := mappers.NewDifficultySvc()
	return services.NewDifficultySvc(difficultiesDataSvc, difficultiesMappingSvc)
}

// createSourcesSvc creates a services.SourceOps.
func createSourcesSvc(db logic.DbOperations, endpoints services.EndpointOps) services.SourceOps {
	sourceDataSvc := data.NewSourceSvc(db)
	sourceMappingSvc := mappers.NewSourceSvc()
	return services.NewSourceSvc(sourceDataSvc, sourceMappingSvc, endpoints)
}

// createEndpointsSvc creates a services.EndpointOps.
func createEndpointsSvc(db logic.DbOperations) services.EndpointOps {
	endpointsDataSvc := data.NewEndpointSvc(db)
	endpointsMappingSvc := mappers.NewEndpointSvc()
	return services.NewEndpointSvc(endpointsDataSvc, endpointsMappingSvc)
}

// createReferencesSvc creates a services.ReferenceOps.
func createReferencesSvc(db logic.DbOperations, sources services.SourceOps) services.ReferenceOps {
	referencesDataSvc := data.NewReferenceSvc(db)
	referencesMappingSvc := mappers.NewReferenceSvc()
	return services.NewReferenceSvc(referencesDataSvc, referencesMappingSvc, sources)
}
