package references

// insertReferenceQuery is a SQL query to insert a new reference record in the 'references' table
const insertReferenceQuery = `
	INSERT INTO "references" (internal_id, source_id, category, type, reference)
    VALUES ($1, $2, $3, $4, $5) 
`

// getReferencesQuery is a SQL query to retrieve references from the provided internal IDs.
const getReferencesQuery = `SELECT internal_id, source_id, category, type, reference FROM "references" WHERE internal_id = ANY($1::uuid[])`
