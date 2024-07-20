package references

// insertReferenceQuery is a SQL query to insert a new reference record in the 'references' table
const insertReferenceQuery = `
	INSERT INTO "references" (internal_id, source_id, category, type, reference)
    VALUES ($1, $2, $3, $4, $5) 
`
