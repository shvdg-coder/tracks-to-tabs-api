package sources

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to sources in the database.
type DataOperations interface {
	InsertSources(sources ...*Source)
	InsertSource(source *Source)
}

// DatabaseService is for managing sources.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DataOperations {
	return &DatabaseService{Database: database}
}

// InsertSources inserts multiple sources in the sources table.
func (a *DatabaseService) InsertSources(sources ...*Source) {
	for _, source := range sources {
		a.InsertSource(source)
	}
}

// InsertSource inserts a new source in the sources table.
func (a *DatabaseService) InsertSource(source *Source) {
	_, err := a.Database.DB.Exec(insertSourceQuery, source.ID, source.Name, source.Category)
	if err != nil {
		log.Printf("Failed inserting source with name: '%s': %s", source.Name, err.Error())
	} else {
		log.Printf("Successfully inserted source with name: '%s'", source.Name)
	}
}
