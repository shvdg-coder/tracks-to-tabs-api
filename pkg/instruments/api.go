package instruments

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"

	"log"
)

// API is for managing tabs and instruments.
type API struct {
	Database *logic.DatabaseManager
}

// NewAPI creates a new instance of the API struct.
func NewAPI(database *logic.DatabaseManager) *API {
	return &API{Database: database}
}

// CreateInstrumentsTable creates an instruments table if it doesn't already exist.
func (a *API) CreateInstrumentsTable() {
	_, err := a.Database.DB.Exec(createInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully created the 'instruments' table.")
	}
}

// DropInstrumentsTable drops the instruments table if it exists.
func (a *API) DropInstrumentsTable() {
	_, err := a.Database.DB.Exec(dropInstrumentsTableQuery)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully dropped the 'instruments' table.")
	}
}

// InsertInstrument inserts a new instrument in the instruments table.
func (a *API) InsertInstrument(name string) {
	_, err := a.Database.DB.Exec(insertInstrumentQuery, name)
	if err != nil {
		log.Printf("Failed inserting instrument with Title: '%s': %s", name, err.Error())
	} else {
		log.Printf("Successfully inserted instrument with Title: '%s'", name)
	}
}
