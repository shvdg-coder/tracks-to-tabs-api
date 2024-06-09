package id_references

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

type API struct {
	Database *logic.DatabaseManager
}

func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateIdReferencesTable creates the id_references table if it doesn't already exist.
func (a *API) CreateIdReferencesTable() {
	_, err := a.Database.DB.Exec(createIdReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'id_references' table")
	}
}

// DropIdReferencesTable drops the id_references table if it exists.
func (a *API) DropIdReferencesTable() {
	_, err := a.Database.DB.Exec(dropIdReferencesTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'id_references' table")
	}
}

// InsertIdReference inserts a record into the id_references table.
func (a *API) InsertIdReference(internalSource, internalID, externalSource, externalID string) {
	_, err := a.Database.DB.Exec(insertIdReferenceQuery, internalSource, internalID, externalSource, externalID)
	if err != nil {
		log.Printf("Failed to insert id reference with InternalSource '%s', InternalID '%s', ExternalSource '%s', and ExternalID '%s': %s", internalSource, internalID, externalSource, externalID, err.Error())
	} else {
		log.Printf("Successfully inserted id reference into the 'id_references' table with InternalSource '%s', InternalID '%s', ExternalSource '%s', and ExternalID '%s'", internalSource, internalID, externalSource, externalID)
	}
}
