package queries

// InsertTab is a SQL query string used to insert a tab into the 'tabs' table.
const InsertTab = `
	INSERT INTO tabs (id, instrument_id, difficulty_id, description)
    VALUES ($1, $2, $3, $4) 
`

// GetTabs is a SQL query sting used to retrieve tabs for the provided IDs.
const GetTabs = `SELECT id, instrument_id, difficulty_id, description FROM tabs WHERE id = ANY($1::uuid[])`
