package data

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// ReferenceData represents operations related to references in the database.
type ReferenceData interface {
	InsertReferenceEntries(reference ...*models.ReferenceEntry) error
	GetReferenceEntries(internalID ...uuid.UUID) ([]*models.ReferenceEntry, error)
	GetReferenceEntry(internalID uuid.UUID) (*models.ReferenceEntry, error)
}

// ReferenceSvc is for managing references.
type ReferenceSvc struct {
	logic.DbOps
}

// NewReferenceSvc creates a new instance of the ReferenceSvc struct.
func NewReferenceSvc(database logic.DbOps) ReferenceData {
	return &ReferenceSvc{DbOps: database}
}

// InsertReferenceEntries inserts multiple references in the references table.
func (d *ReferenceSvc) InsertReferenceEntries(references ...*models.ReferenceEntry) error {
	data := make([][]interface{}, len(references))

	for i, reference := range references {
		data[i] = reference.Fields()
	}

	fieldNames := []string{"internal_id", "source_id", "category", "type", "reference"}
	return d.BulkInsert("references", fieldNames, data)
}

// GetReferenceEntry retrieves a reference entry, without entity references, for the provided internal ID.
func (d *ReferenceSvc) GetReferenceEntry(internalID uuid.UUID) (*models.ReferenceEntry, error) {
	references, err := d.GetReferenceEntries(internalID)
	if err != nil {
		return nil, err
	}
	return references[0], nil
}

// GetReferenceEntries retrieves reference entries, without entity references, for the provided internal IDs.
func (d *ReferenceSvc) GetReferenceEntries(internalID ...uuid.UUID) ([]*models.ReferenceEntry, error) {
	rows, err := d.DB().Query(queries.GetReferences, pq.Array(internalID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var references []*models.ReferenceEntry
	for rows.Next() {
		reference := &models.ReferenceEntry{}
		err := rows.Scan(&reference.InternalID, &reference.SourceID, &reference.Category, &reference.Type, &reference.Reference)
		if err != nil {
			return nil, err
		}
		references = append(references, reference)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return references, nil
}
