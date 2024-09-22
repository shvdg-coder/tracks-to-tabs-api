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

	CREATE INDEX idx_difficulty_id ON difficulties(id);
`

// DropDifficultiesTable is a SQL query to drop the 'difficulties' table if it exists
const DropDifficultiesTable = `
	DROP TABLE IF EXISTS difficulties;
`

// GetDifficulties retrieves the difficulties for the provided IDs.
const GetDifficulties = `SELECT id, name FROM difficulties WHERE id = ANY($1::int[])`
