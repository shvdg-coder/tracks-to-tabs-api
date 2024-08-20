package sources

// insertSourceQuery is a SQL to insert a new record into the 'sources' table
const insertSourceQuery = `
	INSERT INTO sources (id, name, category)
    VALUES ($1, $2, $3) 
`

// getSourcesFromIDs is a SQL query to retrieve source records for the provided ID's
const getSourcesFromIDs = `SELECT id, name, category FROM sources WHERE id = ANY($1::int[])`
