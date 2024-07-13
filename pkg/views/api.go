package views

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	end "github.com/shvdg-dev/tunes-to-tabs-api/pkg/views/endpoints"
	tab "github.com/shvdg-dev/tunes-to-tabs-api/pkg/views/tabs"
)

// API is for managing views.
type API struct {
	Tabs      *tab.API
	Endpoints *end.API
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Tabs: tab.NewAPI(database), Endpoints: end.NewAPI(database)}
}
