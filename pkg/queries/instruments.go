package queries

/*
CreateInstrumentsTable is a query to create an instruments lookup table.
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  distortion guitar  |
| 2  |  acoustic guitar    |
+----+---------------------+

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the instrument.
*/
const CreateInstrumentsTable = `
	CREATE TABLE IF NOT EXISTS instruments (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);

	CREATE INDEX idx_instrument_id ON instruments(id);
`

// DropInstrumentsTable is a SQL query to drop the 'instruments' table if it exists
const DropInstrumentsTable = `
	DROP TABLE IF EXISTS instruments;
`

// GetInstruments is a SQL query string to select an instrument.
const GetInstruments = `SELECT id, name FROM instruments WHERE id = ANY($1::int[])`
