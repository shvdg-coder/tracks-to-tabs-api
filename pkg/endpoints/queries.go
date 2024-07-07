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

/*
+-------------+-------------------+-------------------+-------------------+---------------+--------------------------------------------+
| source_id   |  source_name      | source_category   | endpoint_category | endpoint_type | endpoint_url                               |
+-------------+-------------------+-------------------+-------------------+---------------+--------------------------------------------+
| 1000        | Music Provider 1  | music             | artist            | web           | https://musicprovider1/artist/{artistID}   |
| 1000        | Music Provider 1  | music             | track             | web           | https://musicprovider1/track/{trackID}     |
| 2000        | Tab Provider 1    | tabs              | artist            | web           | https://tabprovider1/artist/{artistID}     |
| 2000        | Tab Provider 1    | tabs              | artist            | api           | https://tabprovider1/artist/api/{artistID} |
+-------------+-------------------+-------------------+-------------------+---------------+--------------------------------------------+

This view is used to display a combination of Source and Endpoint information in our system.

The view combines these tables and includes the following columns:
  - 'source_id': The ID of the source from the `sources` table.
  - 'source_name': The name of the source from the `sources` table.
  - 'source_category': The category of the source from the `sources` table.
  - 'endpoint_category': The category of the endpoint from the `endpoints` table.
  - 'endpoint_type': The type of the endpoint from the `endpoints` table.
  - 'endpoint_url': The URL of the endpoint from the `endpoints` table.
*/
const createSourcesEndpointsViewQuery = `
	CREATE VIEW v_source_endpoints AS
		SELECT sources.id as source_id, sources.name as source_name, sources.category AS source_category, 
			   endpoints.category AS endpoint_category, endpoints.type as endpoint_type, endpoints.url as endpoint_url
		FROM sources
		INNER JOIN endpoints
		ON sources.id = endpoints.source_id;
`

// dropSourcesEndpointsViewQuery is a SQL query to drop the 'sources' to 'endpoints' view from the database.
const dropSourcesEndpointsViewQuery = `
	DROP VIEW IF EXISTS "v_source_endpoints";
`
