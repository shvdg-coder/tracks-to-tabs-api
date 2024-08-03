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

// DataService is for managing sources.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertSources inserts multiple sources in the sources table.
func (d *DataService) InsertSources(sources ...*Source) {
	for _, source := range sources {
		d.InsertSource(source)
	}
}

// InsertSource inserts a new source in the sources table.
func (d *DataService) InsertSource(source *Source) {
	_, err := d.Exec(insertSourceQuery, source.ID, source.Name, source.Category)
	if err != nil {
		log.Printf("Failed inserting source with name: '%s': %s", source.Name, err.Error())
	} else {
		log.Printf("Successfully inserted source with name: '%s'", source.Name)
	}
}
