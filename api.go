package ttt_api

import (
	logic "github.com/shvdg-dev/base-logic/pkg"

	"github.com/shvdg-dev/tunes-to-tabs-api/sessions"
	"github.com/shvdg-dev/tunes-to-tabs-api/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Sessions *sessions.API
	Users    *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{
		Sessions: sessions.NewAPI(database),
		Users:    users.NewAPI(database)}
}
