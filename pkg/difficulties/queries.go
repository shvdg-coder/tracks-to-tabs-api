package difficulties

// insertDifficultyQuery is a SQL query string used to insert a difficulty level into the 'difficulties' table.
const insertDifficultyQuery = `
	INSERT INTO difficulties (name)
    VALUES ($1) 
`
