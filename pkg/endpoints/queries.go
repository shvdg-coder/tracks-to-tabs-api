package endpoints

// insertEndpointQuery is a SQL query to insert a new endpoint record in the 'endpoints' table
const insertEndpointQuery = `
	INSERT INTO "endpoints" (source_id, category, type, url)
    VALUES ($1, $2, $3, $4)
`

// getEndpointsFromIDs is a SQL query to retrieve endpoints from the database.
const getEndpointsFromIDs = `SELECT source_id, category, type, url FROM "endpoints" WHERE source_id = ANY($1::int[])`
