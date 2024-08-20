package references

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	src "github.com/shvdg-dev/tunes-to-tabs-api/pkg/sources"
	"log"
)

// DataOperations represents operations related to references in the database.
type DataOperations interface {
	InsertReference(reference *Reference)
	GetReference(internalID uuid.UUID) (*Reference, error)
	GetReferences(internalID ...uuid.UUID) ([]*Reference, error)
}

// DataService is for managing references.
type DataService struct {
	logic.DbOperations
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database logic.DbOperations) DataOperations {
	return &DataService{DbOperations: database}
}

// InsertReference inserts a record into the references table.
func (d *DataService) InsertReference(reference *Reference) {
	_, err := d.Exec(insertReferenceQuery, reference.InternalID, reference.Source.ID, reference.Category, reference.Type, reference.Reference)
	if err != nil {
		log.Printf(
			"Failed to insert reference with InternalID '%s', SourceID '%d', Category '%s', Type '%s', and Reference '%s': %s",
			reference.InternalID, reference.Source.ID, reference.Category, reference.Type, reference.Reference, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted reference into the 'references' table with InternalID '%s', SourceID '%d', Category '%s', Type '%s', and Reference '%s'",
			reference.InternalID, reference.Source.ID, reference.Category, reference.Type, reference.Reference,
		)
	}
}

// GetReference retrieves a reference
func (d *DataService) GetReference(internalID uuid.UUID) (*Reference, error) {
	references, err := d.GetReferences(internalID)
	if err != nil {
		return nil, err
	}
	return references[0], nil
}

// GetReferences retrieves the references from the provided internal IDs.
func (d *DataService) GetReferences(internalID ...uuid.UUID) ([]*Reference, error) {
	rows, err := d.Query(getReferencesQuery, pq.Array(internalID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var references []*Reference
	for rows.Next() {
		reference := &Reference{}
		source := &src.Source{}
		err := rows.Scan(&reference.InternalID, &source.ID, &reference.Category, &reference.Type, &reference.Reference)
		if err != nil {
			return nil, err
		}
		reference.Source = source
		references = append(references, reference)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return references, nil
}
