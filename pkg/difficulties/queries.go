package difficulties

/*
+----+---------------------+
| ID |         Name        |
+----+---------------------+
| 1  |  Easy               |
| 2  |  Intermediate       |
| 3  |  Hard               |
+----+---------------------+

This table is used to store Difficulty Levels in our system.

It contains the following columns:
  - 'ID': This is an auto-incrementing integer that uniquely identifies a record.
  - 'Name': This column records the name of the difficulty level.
*/
const createDifficultiesTableQuery = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   ID SERIAL PRIMARY KEY,
	   Name VARCHAR(255) NOT NULL
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
