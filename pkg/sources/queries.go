package sources

/*
+------------------------+
| ID |       Name        |
+------------------------+
| 1  | Music Provider 1  |
| 2  | Music Provider 2  |
| 3  | Tab Provider 1    |
| 4  | Tab Provider 2    |
+------------------------+

The table named 'sources' has the purpose of storing unique source names.

It contains the following columns:
  - 'ID': An auto-incrementing integer that uniquely identifies a record.
  - 'Name': The name of the source.
*/
const createSourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS sources(
	   ID int PRIMARY KEY,
	   Name VARCHAR(250) NOT NULL,
	   UNIQUE(Name)                                    
	);
`

// dropSourcesTableQuery is a SQL query to drop the 'sources' table
const dropSourcesTableQuery = `
	DROP TABLE IF EXISTS sources;
`

// insertSourceQuery is a SQL to insert a new record into the 'sources' table
const insertSourceQuery = `
	INSERT INTO sources (ID, Name)
    VALUES ($1, $2) 
`

// getSourcesQuery is a SQL to get all the records from the 'sources' table
const getSourcesQuery = `
	SELECT ID, Name FROM sources
`
