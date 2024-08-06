package references

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to references in the database.
type DataOperations interface {
	InsertReference(reference *Reference)
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
