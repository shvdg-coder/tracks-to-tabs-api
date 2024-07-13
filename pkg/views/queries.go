package views

/*
+------------------+--------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
|  artist_id       | artist_name  |  track_id    |   track_title      |  track_duration  |     tab_id    |   instrument_id   | instrument_name  | difficulty_id  | difficulty_name  | description          |
+---------------------------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
| 234e5678-e90c... | Sting        |     600      | Englishman in NY   |   140305         | 123e4567-e... |  700              |  Bass guitar     | 500            | Intermediate     | Sting playing bass   |
| 234e5649-e90c... | B.B. King    |     650      | The Thrill is Gone |   124022         | 123e4569-e... |  785              |  Electric guitar | 530            | Advanced         | B.B. King's solo     |
+------------------+--------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+

This view is used to display a comprehensive listing of Artist, Track, and Tab information in our system.

The view combines several tables and consists of the following columns:
  - 'artist_id': The ID of the artist from the `artists` table.
  - 'artist_name': The name of the artist from the `artists` table.
  - 'track_id': The ID of the track from the `tracks` table.
  - 'track_title': The title of the track from the `tracks` table.
  - 'track_duration': The duration in milliseconds of the track from the `tracks` table.
  - 'tab_id': The ID of the tab from the `tabs` table.
  - 'instrument_id': The ID of the instrument from the `instruments` table.
  - 'instrument_name': The name of the instrument from the `instruments` table.
  - 'difficulty_id': The ID of the difficulty from the `difficulties` table.
  - 'difficulty_name': The name of the difficulty from the `difficulties` table.
  - 'description': The description of the tab from the `tabs` table.
*/
const createArtistsTracksTabsViewQuery = `
	CREATE VIEW v_artists_tracks_tabs AS
	SELECT 
		a.id AS artist_id,
		a.name AS artist_name,
		tr.id AS track_id,
		tr.title AS track_title,
		tr.duration AS track_duration,
		t.id AS tab_id,
		t.instrument_id,
		i.name AS instrument_name,
		t.difficulty_id,
		d.name AS difficulty_name,
		t.description
	FROM artist_track AS at
	INNER JOIN artists AS a ON at.artist_id = a.id
	INNER JOIN track_tab AS tt ON at.track_id = tt.track_id
	INNER JOIN tracks AS tr ON tt.track_id = tr.id
	INNER JOIN tabs AS t ON tt.tab_id = t.id
	INNER JOIN instruments AS i ON t.instrument_id = i.id
	INNER JOIN difficulties AS d ON t.difficulty_id = d.id;
`

// dropArtistsTracksTabsViewQuery is a SQL query to drop the 'artists' 'tracks' 'tabs' view if it exists
const dropArtistsTracksTabsViewQuery = `
	DROP VIEW IF EXISTS "v_artists_tracks_tabs";
`

// selectArtistsTracksTabsQuery is a SQL query string used to retrieve the artists, tracks and tabs, given a set of IDs, from the database.
const selectArtistsTracksTabsQuery = `
	SELECT * FROM v_artists_tracks_tabs WHERE artist_id = ANY($1);
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
