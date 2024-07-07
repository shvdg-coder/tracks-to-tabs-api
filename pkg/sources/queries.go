package sources

/*
+-------+-------------------+-------------+
| id    |       name        |  category   |
+-------+-------------------+-------------+
| 1001  | Music Provider 1  |   music     |
| 1002  | Music Provider 2  |   music     |
| 1003  | Tab Provider 1    |   tabs      |
| 1004  | Tab Provider 2    |   tabs      |
+-------+-------------------+-------------+

The table named 'sources' has the purpose of storing unique source names.

It contains the following columns:
  - 'id': An auto-incrementing integer that uniquely identifies a record.
  - 'name': The name of the source.
  - 'category': The category of the source.
*/
const createSourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS sources(
	   id int PRIMARY KEY,
	   name VARCHAR(250) NOT NULL,
	   category VARCHAR(100) NOT NULL,
	   UNIQUE(name)                                    
	);
`

// dropSourcesTableQuery is a SQL query to drop the 'sources' table
const dropSourcesTableQuery = `
	DROP TABLE IF EXISTS sources;
`

// insertSourceQuery is a SQL to insert a new record into the 'sources' table
const insertSourceQuery = `
	INSERT INTO sources (id, name, category)
    VALUES ($1, $2, $3) 
`

// getSourcesQuery is a SQL to get all the records from the 'sources' table
const getSourcesQuery = `
	SELECT id, name, category FROM sources
`
