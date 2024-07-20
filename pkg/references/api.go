package references

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// API is for managing references.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// InsertReference inserts a record into the references table.
func (a *API) InsertReference(reference *Reference) {
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
