package queries

// InsertSource is a SQL to insert a new record into the 'sources' table
const InsertSource = `
	INSERT INTO sources (id, name, category)
    VALUES ($1, $2, $3) 
`

// GetSourcesFromIDs is a SQL query to retrieve source records for the provided ID's
const GetSourcesFromIDs = `SELECT id, name, category FROM sources WHERE id = ANY($1::int[])`
