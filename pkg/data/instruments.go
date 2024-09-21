package data

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
	"log"
)

// InstrumentData represents operations related to instruments in the database.
type InstrumentData interface {
	InsertInstrumentEntries(instruments ...*models.InstrumentEntry)
	InsertInstrumentEntry(instrument *models.InstrumentEntry)
	GetInstrumentEntries(instrumentID ...uint) ([]*models.InstrumentEntry, error)
	GetInstrumentEntry(instrumentID uint) (*models.InstrumentEntry, error)
}

// InstrumentSvc is for managing instruments.
type InstrumentSvc struct {
	logic.DbOps
}

// NewInstrumentSvc creates a new instance of the InstrumentSvc struct.
func NewInstrumentSvc(database logic.DbOps) InstrumentData {
	return &InstrumentSvc{DbOperations: database}
}

// InsertInstrumentEntries inserts multiple instruments in the instruments table.
func (d *InstrumentSvc) InsertInstrumentEntries(instruments ...*models.InstrumentEntry) {
	for _, instrument := range instruments {
		d.InsertInstrumentEntry(instrument)
	}
}

// InsertInstrumentEntry inserts a new instrument in the instruments table.
func (d *InstrumentSvc) InsertInstrumentEntry(instrument *models.InstrumentEntry) {
	_, err := d.Exec(queries.InsertInstrument, instrument.Name)
	if err != nil {
		log.Printf("Failed inserting instrument: %s", err.Error())
	}
}

// GetInstrumentEntry retrieves an instrument for the provided ID.
func (d *InstrumentSvc) GetInstrumentEntry(instrumentID uint) (*models.InstrumentEntry, error) {
	instruments, err := d.GetInstrumentEntries(instrumentID)
	if err != nil {
		return nil, err
	}
	return instruments[0], nil
}

// GetInstrumentEntries retrieves instruments for the provided IDs.
func (d *InstrumentSvc) GetInstrumentEntries(instrumentID ...uint) ([]*models.InstrumentEntry, error) {
	rows, err := d.DB().Query(queries.GetInstruments, pq.Array(instrumentID))
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
