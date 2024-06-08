package ttt_api

import (
	"github.com/shvdg-dev/base-pkg/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Sessions *sessions.API
	Users    *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *database.Manager) *API {
	return &API{
		Sessions: sessions.NewAPI(database),
		Users:    users.NewAPI(database)}
}
