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

// DataService is for managing instruments.
type DataService struct {
	*logic.DatabaseManager
}

// NewDataService creates a new instance of the DataService struct.
func NewDataService(database *logic.DatabaseManager) DataOperations {
	return &DataService{DatabaseManager: database}
}

// InsertInstruments inserts multiple instruments in the instruments table.
func (d *DataService) InsertInstruments(instruments ...*Instrument) {
	for _, instrument := range instruments {
		d.InsertInstrument(instrument)
	}
}

// InsertInstrument inserts a new instrument in the instruments table.
func (d *DataService) InsertInstrument(instrument *Instrument) {
	_, err := d.DB.Exec(insertInstrumentQuery, instrument.Name)
	if err != nil {
		log.Printf("Failed inserting instrument with name: '%s': %s", instrument.Name, err.Error())
	} else {
		log.Printf("Successfully inserted instrument with name: '%s'", instrument.Name)
	}
}

// GetInstrument retrieves an instrument for the provided ID.
func (d *DataService) GetInstrument(instrumentID string) (*Instrument, error) {
	instruments, err := d.GetInstruments(instrumentID)
	if err != nil {
		return nil, err
	}
	return instruments[0], nil
}

// GetInstruments retrieves instruments for the provided IDs.
func (d *DataService) GetInstruments(instrumentID ...string) ([]*Instrument, error) {
	rows, err := d.DB.Query(getInstrumentsQuery, instrumentID)
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
