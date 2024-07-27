package difficulties

/*
CreateDifficultiesTableQuery is a query to create a difficulties lookup table.
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
const CreateDifficultiesTableQuery = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

// DropDifficultiesTableQuery is a SQL query to drop the 'difficulties' table if it exists
const DropDifficultiesTableQuery = `
	DROP TABLE IF EXISTS difficulties;
`
