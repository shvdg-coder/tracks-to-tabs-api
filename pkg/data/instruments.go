package data

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/queries"
)

// InstrumentData represents operations related to instruments in the database.
type InstrumentData interface {
	InsertInstrumentEntries(instruments ...*models.InstrumentEntry) error
	GetInstrumentEntries(instrumentID ...uint) ([]*models.InstrumentEntry, error)
	GetInstrumentEntry(instrumentID uint) (*models.InstrumentEntry, error)
}

// InstrumentSvc is for managing instruments.
type InstrumentSvc struct {
	logic.DbOps
}

// NewInstrumentSvc creates a new instance of the InstrumentSvc struct.
func NewInstrumentSvc(database logic.DbOps) InstrumentData {
	return &InstrumentSvc{DbOps: database}
}

// InsertInstrumentEntries inserts multiple instruments in the instruments table.
func (d *InstrumentSvc) InsertInstrumentEntries(instruments ...*models.InstrumentEntry) error {
	data := make([][]interface{}, len(instruments))

	for i, instrument := range instruments {
		data[i] = instrument.Fields()
	}

	fieldNames := []string{"id", "name"}
	return d.BulkInsert("instruments", fieldNames, data)
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
