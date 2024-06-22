package endpoints

/*
+-------------------------------------------------------------------------------------+
|   SourceID   | Category   | Endpoint
+-------------------------------------------------------------------------------------+
| 1001         | Artist     | /artist/$s
| 1002         | Track      | /track/$s
| 1003         | Tab        | /tab/$s
+-------------------------------------------------------------------------------------+

The table 'endpoints' is used to store various endpoints for internal records.

It contains the following columns:
  - 'SourceID': This is the ID of the external source from which the data was referenced.
  - 'Category': This denotes the category of an external reference (e.g., 'Artist', 'Track', 'Tab').
  - 'URL': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const createEndpointsTableQuery = `
	CREATE TABLE IF NOT EXISTS "endpoints" (
	   SourceID INT NOT NULL,
	   Category VARCHAR(250) NOT NULL, 
	   URL VARCHAR(250) NOT NULL,
	   UNIQUE(SourceID, Category, URL),  
	   CONSTRAINT fk_source FOREIGN KEY(SourceID) REFERENCES sources(ID)
	);
`

// dropEndpointsTableQuery is a SQL query to drop the 'endpoints' table from the database.
const dropEndpointsTableQuery = `
	DROP TABLE IF EXISTS "endpoints";
`

// insertEndpointQuery is a SQL query to insert a new endpoint record in the 'endpoints' table
const insertEndpointQuery = `
	INSERT INTO "endpoints" (SourceID, Category, URL)
    VALUES ($1, $2, $3)
`
