package sources

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing sources.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateSourcesTable creates a sources table if it doesn't already exist.
func (a *API) CreateSourcesTable() {
	_, err := a.Database.DB.Exec(createSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' table.")
	}
}

// DropSourcesTable drops the sources table if it exists.
func (a *API) DropSourcesTable() {
	_, err := a.Database.DB.Exec(dropSourcesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' table.")
	}
}

// InsertSources inserts multiple sources in the sources table.
func (a *API) InsertSources(sources ...*Source) {
	for _, source := range sources {
		a.InsertSource(source)
	}
}

// InsertSource inserts a new source in the sources table.
func (a *API) InsertSource(source *Source) {
	_, err := a.Database.DB.Exec(insertSourceQuery, source.ID, source.Name, source.Category)
	if err != nil {
		log.Printf("Failed inserting source with name: '%s': %s", source.Name, err.Error())
	} else {
		log.Printf("Successfully inserted source with name: '%s'", source.Name)
	}
}

// CreateSourcesToEndpointsView creates the sources to endpoints view.
func (a *API) CreateSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(createSourcesToEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'sources' to 'endpoints' view")
	}
}

// DropSourcesToEndpointsView drops the sources to endpoints view if it exists.
func (a *API) DropSourcesToEndpointsView() {
	_, err := a.Database.DB.Exec(dropSourcesToEndpointsViewQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'sources' to 'endpoints' view")
	}
}
