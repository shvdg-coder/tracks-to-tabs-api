package queries

// InsertEndpoint is a SQL query to insert a new endpoint record in the 'endpoints' table
const InsertEndpoint = `
	INSERT INTO "endpoints" (source_id, category, type, url)
    VALUES ($1, $2, $3, $4)
`

// GetEndpointsFromIDs is a SQL query to retrieve endpoints from the database.
const GetEndpointsFromIDs = `SELECT source_id, category, type, url FROM "endpoints" WHERE source_id = ANY($1::int[])`
