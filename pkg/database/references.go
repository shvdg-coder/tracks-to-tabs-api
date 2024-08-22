package database

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// ReferenceOps represents operations related to references in the database.
type ReferenceOps interface {
	InsertReference(reference *models.Reference)
	GetReference(internalID uuid.UUID) (*models.Reference, error)
	GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error)
}

// ReferenceSvc is for managing references.
type ReferenceSvc struct {
	logic.DbOperations
}

// NewReferenceSvc creates a new instance of the ReferenceSvc struct.
func NewReferenceSvc(database logic.DbOperations) ReferenceOps {
	return &ReferenceSvc{DbOperations: database}
}

// InsertReference inserts a record into the references table.
func (d *ReferenceSvc) InsertReference(reference *models.Reference) {
	_, err := d.Exec(queries.InsertReference, reference.InternalID, reference.Source.ID, reference.Category, reference.Type, reference.Reference)
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
func (d *ReferenceSvc) GetReference(internalID uuid.UUID) (*models.Reference, error) {
	references, err := d.GetReferences(internalID)
	if err != nil {
		return nil, err
	}
	return references[0], nil
}

// GetReferences retrieves the references from the provided internal IDs.
func (d *ReferenceSvc) GetReferences(internalID ...uuid.UUID) ([]*models.Reference, error) {
	rows, err := d.Query(queries.GetReferences, pq.Array(internalID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var references []*models.Reference
	for rows.Next() {
		reference := &models.Reference{}
		source := &models.Source{}
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
