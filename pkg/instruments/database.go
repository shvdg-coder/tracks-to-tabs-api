package instruments

import (
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"log"
)

// DataOperations represents operations related to instruments in the database.
type DataOperations interface {
	InsertInstruments(instruments ...*Instrument)
	InsertInstrument(instrument *Instrument)
	GetInstrument(instrumentID string) (*Instrument, error)
	GetInstruments(instrumentID ...string) ([]*Instrument, error)
}

// DatabaseService is for managing instruments.
type DatabaseService struct {
	Database *logic.DatabaseManager
}

// NewDatabaseService creates a new instance of the DatabaseService struct.
func NewDatabaseService(database *logic.DatabaseManager) DataOperations {
	return &DatabaseService{Database: database}
}

// InsertInstruments inserts multiple instruments in the instruments table.
func (a *DatabaseService) InsertInstruments(instruments ...*Instrument) {
	for _, instrument := range instruments {
		a.InsertInstrument(instrument)
	}
}

// InsertInstrument inserts a new instrument in the instruments table.
func (a *DatabaseService) InsertInstrument(instrument *Instrument) {
	_, err := a.Database.DB.Exec(insertInstrumentQuery, instrument.Name)
	if err != nil {
		log.Printf("Failed inserting instrument with name: '%s': %s", instrument.Name, err.Error())
	} else {
		log.Printf("Successfully inserted instrument with name: '%s'", instrument.Name)
	}
}

// GetInstrument retrieves an instrument for the provided ID.
func (a *DatabaseService) GetInstrument(instrumentID string) (*Instrument, error) {
	instruments, err := a.GetInstruments(instrumentID)
	if err != nil {
		return nil, err
	}
	return instruments[0], nil
}

// GetInstruments retrieves instruments for the provided IDs.
func (a *DatabaseService) GetInstruments(instrumentID ...string) ([]*Instrument, error) {
	rows, err := a.Database.DB.Query(getInstrumentsQuery, instrumentID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var instruments []*Instrument
	for rows.Next() {
		var instrument Instrument
		err := rows.Scan(&instrument.ID, &instrument.Name)
		if err != nil {
			return nil, err
		}
		instruments = append(instruments, &instrument)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return instruments, nil
}
