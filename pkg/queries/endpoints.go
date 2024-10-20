package queries

/*
CreateEndpointsTable is a query to create the endpoints table.
It is used to store endpoints, taken from external sources.
+---------------------------------------------------------------+
|   source_id  | category   | type      | url                   |
+---------------------------------------------------------------+
| 1001         | artist     | web       | /artist/{artistID}    |
| 1001         | track      | web       | /track/{trackID}      |
| 1003         | tab        | api       | /tab/{trackID}        |
+---------------------------------------------------------------+

It contains the following columns:
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type.
  - 'url': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const CreateEndpointsTable = `
	CREATE TABLE IF NOT EXISTS "endpoints" (
	   source_id INT REFERENCES sources(id) NOT NULL,
	   category VARCHAR(250) NOT NULL,
	   type VARCHAR(250) NOT NULL,
	   url VARCHAR(250) NOT NULL,
	   PRIMARY KEY (source_id, category, type, url),  
	   CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
	);

	CREATE INDEX idx_endpoint_id ON endpoints(source_id);
`

// DropEndpointsTable is a SQL query to drop the 'endpoints' table from the database.
const DropEndpointsTable = `
	DROP TABLE IF EXISTS "endpoints";
`

// GetEndpointsFromIDs is a SQL query to retrieve endpoints from the database.
const GetEndpointsFromIDs = `SELECT source_id, category, type, url FROM "endpoints" WHERE source_id = ANY($1::int[])`
