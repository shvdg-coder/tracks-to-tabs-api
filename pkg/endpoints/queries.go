package endpoints

/*
+---------------------------------------------------------------+
|   source_id  | category   | type      | url                   |
+---------------------------------------------------------------+
| 1001         | artist     | web       | /artist/{artistID}    |
| 1001         | track      | web       | /track/{trackID}      |
| 1003         | tab        | api       | /tab/{trackID}        |
+---------------------------------------------------------------+

The table 'endpoints' is used to store various endpoints for internal records.

It contains the following columns:
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type.
  - 'url': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const createEndpointsTableQuery = `
	CREATE TABLE IF NOT EXISTS "endpoints" (
	   source_id INT NOT NULL,
	   category VARCHAR(250) NOT NULL,
	   type VARCHAR(250) NOT NULL,
	   url VARCHAR(250) NOT NULL,
	   UNIQUE(source_id, category, type, URL),  
	   CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
	);
`

// dropEndpointsTableQuery is a SQL query to drop the 'endpoints' table from the database.
const dropEndpointsTableQuery = `
	DROP TABLE IF EXISTS "endpoints";
`

// insertEndpointQuery is a SQL query to insert a new endpoint record in the 'endpoints' table
const insertEndpointQuery = `
	INSERT INTO "endpoints" (source_id, category, type, url)
    VALUES ($1, $2, $3, $4)
`
