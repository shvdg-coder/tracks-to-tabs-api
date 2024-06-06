package ttt_api

import (
	"github.com/shvdg-dev/base-pkg/database"
	"github.com/shvdg-dev/ttt-api/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Users *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(db *database.Manager) *API {
	return &API{Users: users.NewAPI(db)}
}
