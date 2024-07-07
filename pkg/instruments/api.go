package instruments

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"

	"log"
)

// API is for managing instruments.
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

// InsertInstruments inserts multiple instruments in the instruments table.
func (a *API) InsertInstruments(instruments ...*Instrument) {
	for _, instrument := range instruments {
		a.InsertInstrument(instrument)
	}
}

// InsertInstrument inserts a new instrument in the instruments table.
func (a *API) InsertInstrument(instrument *Instrument) {
	_, err := a.Database.DB.Exec(insertInstrumentQuery, instrument.Name)
	if err != nil {
		log.Printf("Failed inserting instrument with name: '%s': %s", instrument.Name, err.Error())
	} else {
		log.Printf("Successfully inserted instrument with name: '%s'", instrument.Name)
	}
}
