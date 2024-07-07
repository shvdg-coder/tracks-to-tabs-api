package endpoints

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing endpoints.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateEndpointsTable creates the endpoints table if it doesn't already exist.
func (a *API) CreateEndpointsTable() {
	_, err := a.Database.DB.Exec(createEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'endpoints' table")
	}
}

// DropEndpointsTable drops the endpoints table if it exists.
func (a *API) DropEndpointsTable() {
	_, err := a.Database.DB.Exec(dropEndpointsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'endpoints' table")
	}
}

// InsertEndpoints inserts multiple records into the endpoints table.
func (a *API) InsertEndpoints(endpoints ...*Endpoint) {
	for _, endpoint := range endpoints {
		a.InsertEndpoint(endpoint)
	}
}

// InsertEndpoint inserts a record into the endpoints table.
func (a *API) InsertEndpoint(endpoint *Endpoint) {
	_, err := a.Database.DB.Exec(insertEndpointQuery, endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL)
	if err != nil {
		log.Printf(
			"Failed to insert endpoint with SourceID '%d', Category '%s', Type '%s', and URL '%s': %s",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted endpoint into the 'endpoints' table with SourceID '%d', Category '%s', Type '%s', and URL '%s'",
			endpoint.SourceID, endpoint.Category, endpoint.Type, endpoint.URL,
		)
	}
}
