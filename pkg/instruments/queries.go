package instruments

// insertInstrumentQuery is a SQL query string used to insert an instrument into the 'instruments' table.
const insertInstrumentQuery = `
	INSERT INTO instruments (name)
    VALUES ($1) 
`

// getInstrumentsQuery is a SQL query string to select an instrument.
const getInstrumentsQuery = `SELECT id, name FROM instruments WHERE id IN ($1)`
