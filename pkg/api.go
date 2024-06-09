package pkg

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/artists"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/id_references"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Artists      *artists.API
	IdReferences *id_references.API
	Sessions     *sessions.API
	Users        *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{
		Artists:      artists.NewAPI(database),
		IdReferences: id_references.NewAPI(database),
		Sessions:     sessions.NewAPI(database),
		Users:        users.NewAPI(database),
	}
}
