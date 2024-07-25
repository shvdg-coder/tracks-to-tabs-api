package references

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DatabaseOperations represents operations related to references in the database.
type DatabaseOperations interface {
	InsertReference(reference *Reference)
}

// DatabaseService is for managing references.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DatabaseOperations {
	return &DatabaseService{Database: database}
}

// InsertReference inserts a record into the references table.
func (a *DatabaseService) InsertReference(reference *Reference) {
	_, err := a.Database.DB.Exec(insertReferenceQuery, reference.InternalID, reference.SourceID, reference.Category, reference.Type, reference.Reference)
	if err != nil {
		log.Printf(
			"Failed to insert reference with InternalID '%s', SourceID '%s', Category '%s', Type '%s', and Reference '%s': %s",
			reference.InternalID, reference.SourceID, reference.Category, reference.Type, reference.Reference, err.Error(),
		)
	} else {
		log.Printf(
			"Successfully inserted reference into the 'references' table with InternalID '%s', SourceID '%s', Category '%s', Type '%s', and Reference '%s'",
			reference.InternalID, reference.SourceID, reference.Category, reference.Type, reference.Reference,
		)
	}
}
