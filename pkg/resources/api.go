package resources

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing resources.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateResourcesTable creates the resources table if it doesn't already exist.
func (a *API) CreateResourcesTable() {
	_, err := a.Database.DB.Exec(createResourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'resources' table")
	}
}

// DropResourcesTable drops the resources table if it exists.
func (a *API) DropResourcesTable() {
	_, err := a.Database.DB.Exec(dropResourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'resources' table")
	}
}

// InsertResource inserts a record into the resources table.
func (a *API) InsertResource(resource *Resource) {
	_, err := a.Database.DB.Exec(insertResourceQuery, resource.InternalID, resource.SourceID, resource.Category, resource.Type, resource.Resource)
	if err != nil {
		log.Printf(
			"Failed to insert resource with InternalID '%s', SourceID '%s', Category '%s', Type '%s', and Resource '%s': %s",
			resource.InternalID, resource.SourceID, resource.Category, resource.Type, resource.Resource, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted resource into the 'resources' table with InternalID '%s', SourceID '%s', Category '%s', Type '%s', and Resource '%s'",
			resource.InternalID, resource.SourceID, resource.Category, resource.Type, resource.Resource,
		)
	}
}
