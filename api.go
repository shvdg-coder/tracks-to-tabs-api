package ttt_api

import (
	"tabs/pkg/base/database"
	"tabs/pkg/ttt-api/users"
)

// API represents the main entry point to interact with the API functionalities.
type API struct {
	Users *users.API
}

// NewAPI creates a new instance of the API.
func NewAPI(db *database.Manager) *API {
	return &API{Users: users.NewAPI(db)}
}
