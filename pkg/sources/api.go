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
