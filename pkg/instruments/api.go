package instruments

import (
	"database/sql"
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

// GetInstruments retrieves the instruments.
func (a *API) GetInstruments() []*Instrument {
	rows, err := a.Database.DB.Query(getInstrumentsQuery)
	if err != nil {
		log.Printf("Failed to get instruments: %s", err)
		return nil
	}

	instruments := rowsToInstruments(rows)
	defer closeRows(rows)

	return instruments
}

// rowsToInstruments converts the given *sql.Rows into a slice of *Instrument objects.
func rowsToInstruments(rows *sql.Rows) []*Instrument {
	var instruments []*Instrument
	for rows.Next() {
		instrument := rowsToInstrument(rows)
		if instrument != nil {
			instruments = append(instruments, instrument)
		}
	}
	return instruments
}

// rowsToInstrument scans the SQL row into an Instrument struct.
func rowsToInstrument(rows *sql.Rows) *Instrument {
	var instrument Instrument
	err := rows.Scan(&instrument.ID, &instrument.Name)
	if err != nil {
		log.Printf("Unable to scan instrument: %s", err.Error())
		return nil
	}
	return &instrument
}

// closeRows closes the SQL rows and logs error if any.
func closeRows(rows *sql.Rows) {
	err := rows.Err()
	if err != nil {
		log.Printf("Error while processing rows: %s", err.Error())
	}
	err = rows.Close()
	if err != nil {
		log.Printf("Failed to close rows: %s", err.Error())
	}
}
