package instruments

// insertInstrumentQuery is a SQL query string used to insert an instrument into the 'instruments' table.
const insertInstrumentQuery = `
	INSERT INTO instruments (name)
    VALUES ($1) 
`
