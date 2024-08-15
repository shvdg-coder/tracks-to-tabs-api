package pkg

import (
	"github.com/google/uuid"
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
	GetArtistsCascading(artistID ...uuid.UUID) ([]*art.Artist, error)
	trk.DataOperations
	GetTracksCascading(tabID ...uuid.UUID) ([]*trk.Track, error)
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
	Operations
}

// NewAPI creates a new instance of the API.
func NewAPI(database logic.DbOperations) DataOperations {
	return &API{Operations: NewServiceManager(database)}
}
