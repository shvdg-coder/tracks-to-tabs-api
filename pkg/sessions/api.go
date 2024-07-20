package sessions

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
)

// API is for managing users.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}
