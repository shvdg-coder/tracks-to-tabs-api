package views

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
	CREATE VIEW v_sources_endpoints AS
		SELECT sources.id as source_id, sources.name as source_name, sources.category AS source_category, 
			   endpoints.category AS endpoint_category, endpoints.type as endpoint_type, endpoints.url as endpoint_url
		FROM sources
		INNER JOIN endpoints
		ON sources.id = endpoints.source_id;
`

// dropSourcesEndpointsViewQuery is a SQL query to drop the 'sources and endpoints' view from the database.
const dropSourcesEndpointsViewQuery = `
	DROP VIEW IF EXISTS "v_sources_endpoints";
`

// selectSourcesEndpoints is a SQL query used to retrieve all the sources and endpoints from the database.
const selectSourcesEndpoints = `
	SELECT * FROM v_sources_endpoints;
`
