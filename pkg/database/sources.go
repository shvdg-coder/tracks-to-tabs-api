package database

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// SourceOps represents operations related to sources in the database.
type SourceOps interface {
	InsertSources(sources ...*models.Source)
	InsertSource(source *models.Source)
	GetSource(id uint) (*models.Source, error)
	GetSources(id ...uint) ([]*models.Source, error)
}

// SourceSvc is for managing sources.
type SourceSvc struct {
	logic.DbOperations
}

// NewSourceSvc creates a new instance of the SourceSvc struct.
func NewSourceSvc(database logic.DbOperations) SourceOps {
	return &SourceSvc{DbOperations: database}
}

// InsertSources inserts multiple sources in the sources table.
func (d *SourceSvc) InsertSources(sources ...*models.Source) {
	for _, source := range sources {
		d.InsertSource(source)
	}
}

// InsertSource inserts a new source in the sources table.
func (d *SourceSvc) InsertSource(source *models.Source) {
	_, err := d.Exec(queries.InsertSource, source.ID, source.Name, source.Category)
	if err != nil {
		log.Printf("Failed inserting source with name: '%s': %s", source.Name, err.Error())
	} else {
		log.Printf("Successfully inserted source with name: '%s'", source.Name)
	}
}

// GetSource retrieves a source from the database.
func (d *SourceSvc) GetSource(id uint) (*models.Source, error) {
	sources, err := d.GetSources(id)
	if err != nil {
		return nil, err
	}
	return sources[0], nil
}

// GetSources retrieves multiple sources from the database.
func (d *SourceSvc) GetSources(sourceID ...uint) ([]*models.Source, error) {
	rows, err := d.Query(queries.GetSourcesFromIDs, pq.Array(sourceID))
	if err != nil {
		return nil, err
	}

	var sources []*models.Source

	for rows.Next() {
		source := &models.Source{}
		source.Endpoints = make([]*models.EndpointEntry, 0)
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
