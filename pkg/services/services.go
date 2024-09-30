package services

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/data"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/schemas"
)

// SvcOps represents operations regarding all the services.
type SvcOps interface {
	ArtistOps
	TrackOps
	TabOps
	InstrumentOps
	DifficultyOps
	ReferenceOps
	SourceOps
	EndpointOps
	ArtistTrackOps
	TrackTabOps
	UserOps
	ResourceOps
}

// SvcManager instantiates and handles the different services.
type SvcManager struct {
	ArtistOps
	TrackOps
	TabOps
	InstrumentOps
	DifficultyOps
	ReferenceOps
	SourceOps
	EndpointOps
	ArtistTrackOps
	TrackTabOps
	UserOps
	ResourceOps
}

// NewSvcManager creates a new instance of the SvcManager.
func NewSvcManager(database logic.DbOps) *SvcManager {
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
func createArtistsSvc(db logic.DbOps, artistTrack ArtistTrackOps, tracks TrackOps, references ReferenceOps) ArtistOps {
	artistDataSvc := data.NewArtistSvc(db)
	artistSchemaSvc := schemas.NewArtistSvc(db)
	artistMappingSvc := mappers.NewArtistSvc()
	return NewArtistSvc(artistSchemaSvc, artistDataSvc, artistMappingSvc, artistTrack, tracks, references)
}

// createTracksSvc creates a services.TrackOps.
func createTracksSvc(db logic.DbOps, trackTab TrackTabOps, tabs TabOps, references ReferenceOps) TrackOps {
	trackSchemaSvc := schemas.NewTrackSvc(db)
	trackDataSvc := data.NewTrackSvc(db)
	trackMappingSvc := mappers.NewTrackSvc()
	return NewTrackSvc(trackSchemaSvc, trackDataSvc, trackMappingSvc, trackTab, tabs, references)
}

// createTabsSvc creates a services.TabOps.
func createTabsSvc(db logic.DbOps, instruments InstrumentOps, difficulties DifficultyOps, references ReferenceOps) TabOps {
	tabDataSvc := data.NewTabSvc(db)
	tabSchemaSvc := schemas.NewTabSvc(db)
	tabMappingSvc := mappers.NewTabSvc()
	return NewTabSvc(tabSchemaSvc, tabDataSvc, tabMappingSvc, instruments, difficulties, references)
}

// createArtistTrackService creates an services.ArtistTrackOps.
func createArtistTrackService(db logic.DbOps) ArtistTrackOps {
	artistTrackSchemaSvc := schemas.NewArtistTrackSvc(db)
	artistTrackDataSvc := data.NewArtistTrackSvc(db)
	return NewArtistTrackSvc(artistTrackSchemaSvc, artistTrackDataSvc)
}

// createTrackTabSvc creates a services.TrackTabOps.
func createTrackTabSvc(db logic.DbOps) TrackTabOps {
	trackTabSchemaSvc := schemas.NewTrackTabSvc(db)
	trackTabDataSvc := data.NewTrackTabSvc(db)
	return NewTrackTabSvc(trackTabSchemaSvc, trackTabDataSvc)
}

// createUsersService creates a services.UserOps.
func createUsersService(db logic.DbOps) UserOps {
	userSchemaSvc := schemas.NewUserSvc(db)
	usersDataSvc := data.NewUserSvc(db)
	return NewUserSvc(userSchemaSvc, usersDataSvc)
}

// createInstrumentsSvc creates a services.InstrumentOps.
func createInstrumentsSvc(db logic.DbOps) InstrumentOps {
	instrumentSchemaSvc := schemas.NewInstrumentSvc(db)
	instrumentsDataSvc := data.NewInstrumentSvc(db)
	instrumentMappingSvc := mappers.NewInstrumentSvc()
	return NewInstrumentSvc(instrumentSchemaSvc, instrumentsDataSvc, instrumentMappingSvc)
}

// createDifficultiesSvc creates a services.DifficultyOps.
func createDifficultiesSvc(db logic.DbOps) DifficultyOps {
	difficultySchemaSvc := schemas.NewDifficultySvc(db)
	difficultiesDataSvc := data.NewDifficultySvc(db)
	difficultiesMappingSvc := mappers.NewDifficultySvc()
	return NewDifficultySvc(difficultySchemaSvc, difficultiesDataSvc, difficultiesMappingSvc)
}

// createSourcesSvc creates a services.SourceOps.
func createSourcesSvc(db logic.DbOps, endpoints EndpointOps) SourceOps {
	sourceSchemaSvc := schemas.NewSourceSvc(db)
	sourceDataSvc := data.NewSourceSvc(db)
	sourceMappingSvc := mappers.NewSourceSvc()
	return NewSourceSvc(sourceSchemaSvc, sourceDataSvc, sourceMappingSvc, endpoints)
}

// createEndpointsSvc creates a services.EndpointOps.
func createEndpointsSvc(db logic.DbOps) EndpointOps {
	endpointSchemaSvc := schemas.NewEndpointSvc(db)
	endpointsDataSvc := data.NewEndpointSvc(db)
	endpointsMappingSvc := mappers.NewEndpointSvc()
	return NewEndpointSvc(endpointSchemaSvc, endpointsDataSvc, endpointsMappingSvc)
}

// createReferencesSvc creates a services.ReferenceOps.
func createReferencesSvc(db logic.DbOps, sources SourceOps) ReferenceOps {
	referenceSchemaSvc := schemas.NewReferenceSvc(db)
	referencesDataSvc := data.NewReferenceSvc(db)
	referencesMappingSvc := mappers.NewReferenceSvc()
	return NewReferenceSvc(referenceSchemaSvc, referencesDataSvc, referencesMappingSvc, sources)
}

// createResourcesSvc creates a services.ResourceSvc.
func createResourcesSvc() ResourceOps {
	return NewResourceSvc()
}
