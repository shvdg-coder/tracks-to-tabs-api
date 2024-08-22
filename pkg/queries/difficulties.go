package queries

// InsertDifficulty is a SQL query string used to insert a difficulty level into the 'difficulties' table.
const InsertDifficulty = `
	INSERT INTO difficulties (name)
    VALUES ($1) 
`

// GetDifficulties retrieves the difficulties for the provided IDs.
const GetDifficulties = `SELECT id, name FROM difficulties WHERE id = ANY($1::int[])`
