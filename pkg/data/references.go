package data

import (
	"database/sql"
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// ReferenceData represents operations related to references in the database.
type ReferenceData interface {
	InsertReferenceEntries(reference ...*models.ReferenceEntry) error
	GetReferenceEntries(internalID ...uuid.UUID) ([]*models.ReferenceEntry, error)
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

// GetReferenceEntries retrieves reference entries, without entity references, for the provided internal IDs.
func (d *ReferenceSvc) GetReferenceEntries(internalIDs ...uuid.UUID) ([]*models.ReferenceEntry, error) {
	return logic.BatchGet(d, batchSize, queries.GetReferences, internalIDs, scanReferenceEntry)
}

// scanReferenceEntry scans a row into a models.ReferenceEntry.
func scanReferenceEntry(rows *sql.Rows) (*models.ReferenceEntry, error) {
	reference := &models.ReferenceEntry{}
	if err := rows.Scan(&reference.InternalID, &reference.SourceID, &reference.Category, &reference.Type, &reference.Reference); err != nil {
		return nil, err
	}
	return reference, nil
}
