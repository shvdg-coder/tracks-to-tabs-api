package difficulties

/*
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  easy               |
| 2  |  intermediate       |
| 3  |  hard               |
+----+---------------------+

This table is used to store Difficulty Levels in our system.

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the difficulty level.
*/
const createDifficultiesTableQuery = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

// dropDifficultiesTableQuery is a SQL query to drop the 'difficulties' table if it exists
const dropDifficultiesTableQuery = `
	DROP TABLE IF EXISTS difficulties;
`

// insertDifficultyQuery is a SQL query string used to insert a difficulty level into the 'difficulties' table.
const insertDifficultyQuery = `
	INSERT INTO difficulties (name)
    VALUES ($1) 
`

// getDifficultiesQuery is a SQL query to get difficulties.
const getDifficultiesQuery = `SELECT id, name FROM difficulties`
