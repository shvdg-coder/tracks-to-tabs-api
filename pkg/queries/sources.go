package queries

/*
CreateSourcesTableQuery is a query to create sources table.
It is used to store external sources.
+-------+-------------------+-------------+
| id    |       name        |  category   |
+-------+-------------------+-------------+
| 1001  | Music Provider 1  |   music     |
| 1002  | Music Provider 2  |   music     |
| 1003  | Tab Provider 1    |   tabs      |
| 1004  | Tab Provider 2    |   tabs      |
+-------+-------------------+-------------+

It contains the following columns:
  - 'id': An auto-incrementing integer that uniquely identifies a record.
  - 'name': The name of the source.
  - 'category': The category of the source.
*/
const CreateSourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS sources(
	   id int PRIMARY KEY,
	   name VARCHAR(250) NOT NULL,
	   category VARCHAR(100) NOT NULL,
	   UNIQUE(name)                                    
	);
`

// DropSourcesTableQuery is a SQL query to drop the 'sources' table
const DropSourcesTableQuery = `
	DROP TABLE IF EXISTS sources;
`

// InsertSource is a SQL to insert a new record into the 'sources' table
const InsertSource = `
	INSERT INTO sources (id, name, category)
    VALUES ($1, $2, $3) 
`

// GetSourcesFromIDs is a SQL query to retrieve source records for the provided ID's
const GetSourcesFromIDs = `SELECT id, name, category FROM sources WHERE id = ANY($1::int[])`
