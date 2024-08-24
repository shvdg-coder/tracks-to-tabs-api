package queries

/*
CreateDifficultiesTable is a query to create a difficulties lookup table.
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  easy               |
| 2  |  intermediate       |
| 3  |  hard               |
+----+---------------------+

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the difficulty level.
*/
const CreateDifficultiesTable = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

// DropDifficultiesTable is a SQL query to drop the 'difficulties' table if it exists
const DropDifficultiesTable = `
	DROP TABLE IF EXISTS difficulties;
`

// InsertDifficulty is a SQL query string used to insert a difficulty level into the 'difficulties' table.
const InsertDifficulty = `
	INSERT INTO difficulties (name)
    VALUES ($1) 
`

// GetDifficulties retrieves the difficulties for the provided IDs.
const GetDifficulties = `SELECT id, name FROM difficulties WHERE id = ANY($1::int[])`
