package sources

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to sources in the database.
type DataOperations interface {
	InsertSources(sources ...*Source)
	InsertSource(source *Source)
	GetSource(id uint) (*Source, error)
	GetSources(id ...uint) ([]*Source, error)
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

// GetSource retrieves a source from the database.
func (d *DataService) GetSource(id uint) (*Source, error) {
	sources, err := d.GetSources(id)
	if err != nil {
		return nil, err
	}
	return sources[0], nil
}

// GetSources retrieves multiple sources from the database.
func (d *DataService) GetSources(sourceID ...uint) ([]*Source, error) {
	rows, err := d.Query(getSourcesFromIDs, pq.Array(sourceID))
	if err != nil {
		return nil, err
	}

	var sources []*Source

	for rows.Next() {
		source := &Source{}
		err = rows.Scan(&source.ID, &source.Name, &source.Category)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return sources, nil
}
