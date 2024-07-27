package endpoints

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to endpoints in the database.
type DataOperations interface {
	InsertEndpoints(endpoints ...*Endpoint)
	InsertEndpoint(endpoint *Endpoint)
}

// DatabaseService is for managing endpoints.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of DatabaseService.
func NewDatabaseService(database *logic.DatabaseManager) DataOperations {
	return &DatabaseService{Database: database}
}

// InsertEndpoints inserts multiple records into the endpoints table.
func (a *DatabaseService) InsertEndpoints(endpoints ...*Endpoint) {
	for _, endpoint := range endpoints {
		a.InsertEndpoint(endpoint)
	}
}

// InsertEndpoint inserts a record into the endpoints table.
func (a *DatabaseService) InsertEndpoint(endpoint *Endpoint) {
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
