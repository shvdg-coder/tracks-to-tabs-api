package tabs

// insertTabQuery is a SQL query string used to insert a tab into the 'tabs' table.
const insertTabQuery = `
	INSERT INTO tabs (id, instrument_id, difficulty_id, description)
    VALUES ($1, $2, $3, $4) 
`

// getTabsQuery is a SQL query sting used to retrieve tabs for the provided IDs.
const getTabsQuery = `SELECT id, instrument_id, difficulty_id, description FROM tabs WHERE id = ANY($1::uuid[])`
