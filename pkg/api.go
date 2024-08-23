package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/data"
)

// DataOps represents all API data operations.
type DataOps interface {
	data.ArtistData
	data.TrackData
	data.TabData
	data.DifficultyData
	data.InstrumentData
	data.ReferenceData
	data.SourceData
	data.EndpointsData
	data.ArtistTrackData
	data.TrackTabData
	data.UserData
}

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	SvcOps
}

// NewAPI creates a new instance of the API.
func NewAPI(database logic.DbOperations) DataOps {
	return &API{SvcOps: NewSvcManager(database)}
}
