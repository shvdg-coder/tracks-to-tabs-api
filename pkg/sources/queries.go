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
const createSourcesToEndpointsViewQuery = `
	CREATE VIEW v_sources_to_endpoints AS
		SELECT sources.id as source_id, sources.name as source_name, sources.category AS source_category, 
			   endpoints.category AS endpoint_category, endpoints.type as endpoint_type, endpoints.url as endpoint_url
		FROM sources
		INNER JOIN endpoints
		ON sources.id = endpoints.source_id;
`

// dropSourcesToEndpointsViewQuery is a SQL query to drop the 'sources' to 'endpoints' view from the database.
const dropSourcesToEndpointsViewQuery = `
	DROP VIEW IF EXISTS "v_sources_to_endpoints";
`
