package instruments

/*
+----+---------------------+
| ID |         Name        |
+----+---------------------+
| 1  |  Distortion Guitar  |
| 2  |  Acoustic Guitar    |
+----+---------------------+

This table is used to store Instruments in our system.

It contains the following columns:
  - 'ID': This is an auto-incrementing integer that uniquely identifies a record.
  - 'Name': This column records the name of the instrument.
*/
const createInstrumentsTableQuery = `
	CREATE TABLE IF NOT EXISTS instruments (
	   ID SERIAL PRIMARY KEY,
	   Name VARCHAR(255) NOT NULL
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
	SELECT ID, Name FROM instruments
`
