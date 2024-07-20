package sources

// insertSourceQuery is a SQL to insert a new record into the 'sources' table
const insertSourceQuery = `
	INSERT INTO sources (id, name, category)
    VALUES ($1, $2, $3) 
`
