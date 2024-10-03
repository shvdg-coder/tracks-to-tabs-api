package data

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// SourceData represents operations related to sources in the database.
type SourceData interface {
	InsertSourceEntries(sources ...*models.SourceEntry) error
	GetSourceEntries(id ...uint) ([]*models.SourceEntry, error)
	GetSourceEntry(id uint) (*models.SourceEntry, error)
}

// SourceSvc is for managing sources.
type SourceSvc struct {
	logic.DbOps
}

// NewSourceSvc creates a new instance of the SourceSvc struct.
func NewSourceSvc(database logic.DbOps) SourceData {
	return &SourceSvc{DbOps: database}
}

// InsertSourceEntries inserts multiple sources in the sources table.
func (d *SourceSvc) InsertSourceEntries(sources ...*models.SourceEntry) error {
	data := make([][]interface{}, len(sources))

	for i, source := range sources {
		data[i] = logic.GetFields(source)
	}

	return d.BulkInsert("sources", logic.GetFieldNames("db", &models.SourceEntry{}), data)
}

// GetSourceEntry retrieves a source from the database.
func (d *SourceSvc) GetSourceEntry(id uint) (*models.SourceEntry, error) {
	sources, err := d.GetSourceEntries(id)
	if err != nil {
		return nil, err
	}
	return sources[0], nil
}

// GetSourceEntries retrieves multiple sources from the database.
func (d *SourceSvc) GetSourceEntries(sourceID ...uint) ([]*models.SourceEntry, error) {
	rows, err := d.DB().Query(queries.GetSourcesFromIDs, pq.Array(sourceID))
	if err != nil {
		return nil, err
	}

	var sources []*models.SourceEntry

	for rows.Next() {
		source := &models.SourceEntry{}
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
