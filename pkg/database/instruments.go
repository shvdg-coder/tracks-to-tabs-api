package database

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg/queries"
	"log"
)

// InstrumentOps represents operations related to instruments in the database.
type InstrumentOps interface {
	InsertInstruments(instruments ...*models.InstrumentEntry)
	InsertInstrument(instrument *models.InstrumentEntry)
	GetInstrument(instrumentID uint) (*models.InstrumentEntry, error)
	GetInstruments(instrumentID ...uint) ([]*models.InstrumentEntry, error)
}

// InstrumentSvc is for managing instruments.
type InstrumentSvc struct {
	logic.DbOperations
}

// NewInstrumentSvc creates a new instance of the InstrumentSvc struct.
func NewInstrumentSvc(database logic.DbOperations) InstrumentOps {
	return &InstrumentSvc{DbOperations: database}
}

// InsertInstruments inserts multiple instruments in the instruments table.
func (d *InstrumentSvc) InsertInstruments(instruments ...*models.InstrumentEntry) {
	for _, instrument := range instruments {
		d.InsertInstrument(instrument)
	}
}

// InsertInstrument inserts a new instrument in the instruments table.
func (d *InstrumentSvc) InsertInstrument(instrument *models.InstrumentEntry) {
	_, err := d.Exec(queries.InsertInstrument, instrument.Name)
	if err != nil {
		log.Printf("Failed inserting instrument with name: '%s': %s", instrument.Name, err.Error())
	} else {
		log.Printf("Successfully inserted instrument with name: '%s'", instrument.Name)
	}
}

// GetInstrument retrieves an instrument for the provided ID.
func (d *InstrumentSvc) GetInstrument(instrumentID uint) (*models.InstrumentEntry, error) {
	instruments, err := d.GetInstruments(instrumentID)
	if err != nil {
		return nil, err
	}
	return instruments[0], nil
}

// GetInstruments retrieves instruments for the provided IDs.
func (d *InstrumentSvc) GetInstruments(instrumentID ...uint) ([]*models.InstrumentEntry, error) {
	rows, err := d.Query(queries.GetInstruments, pq.Array(instrumentID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var instruments []*models.InstrumentEntry
	for rows.Next() {
		instrument := &models.InstrumentEntry{}
		err := rows.Scan(&instrument.ID, &instrument.Name)
		if err != nil {
			return nil, err
		}
		instruments = append(instruments, instrument)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return instruments, nil
}
