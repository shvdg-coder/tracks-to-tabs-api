package instruments

/*
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  distortion guitar  |
| 2  |  acoustic guitar    |
+----+---------------------+

This table is used to store Instruments in our system.

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the instrument.
*/
const createInstrumentsTableQuery = `
	CREATE TABLE IF NOT EXISTS instruments (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

// dropInstrumentsTableQuery is a SQL query to drop the 'instruments' table if it exists
const dropInstrumentsTableQuery = `
	DROP TABLE IF EXISTS instruments;
`

// insertInstrumentQuery is a SQL query string used to insert an instrument into the 'instruments' table.
const insertInstrumentQuery = `
	INSERT INTO instruments (name)
    VALUES ($1) 
`

// getInstrumentsQuery is a SQL query string used to retrieve all instruments from the 'instruments' table.
const getInstrumentsQuery = `
	SELECT id, name FROM instruments
`
