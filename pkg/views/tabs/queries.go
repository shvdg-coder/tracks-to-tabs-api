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
