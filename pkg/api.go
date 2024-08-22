package pkg

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
)

// DataOperations represents all API data operations.
type DataOperations interface {
	database.ArtistOps
	GetArtistsCascading(artistID ...uuid.UUID) ([]*models.ArtistEntry, error)
	trk.DataOperations
	GetTracksCascading(tabID ...uuid.UUID) ([]*models.Track, error)
	database.TabsOps
	database.TrackOps
	database.TabsOps
	database.UserOps
	database.TabsOps
	database.TabsOps
	database.TabsOps
	database.TabsOps
	database.TabsOps
}

// API represents the main entry point to interact with functionalities for the defined entities.
type API struct {
	Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database logic.DbOperations) DataOperations {
	return &API{Operations: NewServiceManager(database)}
}
