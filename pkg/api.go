package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/id_references"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/instruments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tabs"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/tracks"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Artists      *artists.API
	IdReferences *id_references.API
	Instruments  *instruments.API
	Sessions     *sessions.API
	Tabs         *tabs.API
	Tracks       *tracks.API
	Users        *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{
		Artists:      artists.NewAPI(database),
		IdReferences: id_references.NewAPI(database),
		Instruments:  instruments.NewAPI(database),
		Sessions:     sessions.NewAPI(database),
		Tabs:         tabs.NewAPI(database),
		Tracks:       tracks.NewAPI(database),
		Users:        users.NewAPI(database),
	}
}
