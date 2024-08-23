package pkg

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
)

// DataOps represents all API data operations.
type DataOps interface {
	GetArtists(artistID ...uuid.UUID) ([]*models.Artist, error)
	GetTracks(trackID ...uuid.UUID) ([]*models.Track, error)
	GetTabs(tabID ...uuid.UUID) ([]*models.Tab, error)
	GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error)
	GetSources(sourceID ...uint) ([]*models.Source, error)
}

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	SvcOps
}

// NewAPI creates a new instance of the API.
func NewAPI(database logic.DbOperations) DataOps {
	return &API{SvcOps: NewSvcManager(database)}
}
