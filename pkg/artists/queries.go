package artists

/*
+--------------------------------------+------------+
|                   id                 |   name    |
+--------------------------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Artist 1   |
| 123e4567-e89b-12d3-a456-426614174001 | Artist 2   |
+--------------------------------------+------------+

This table is used to store Artists in our system.

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record in our system.
  - 'name': This column has the name of an Artist.
*/
const createArtistsTableQuery = `
	CREATE TABLE IF NOT EXISTS artists  (
	   id UUID PRIMARY KEY,
	   name VARCHAR(500) NOT NULL
	);
`

// dropArtistsTableQuery is a SQL query to drop the 'artists' table from the database.
const dropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`

// insertArtistQuery is a SQL query string used to insert an artist into the 'artists' table.
const insertArtistQuery = `
	INSERT INTO artists (id, name)
    VALUES ($1, $2) 
`

/*
+--------------------------------------+--------------------------------------+
|              artist_id               |              track_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is linking the 'artists' table and the 'tracks' table.

- 'artist_id': The UUID that uniquely identifies an artist in the 'artists' table.
- 'track_id': The UUID that uniquely identifies a track in the 'tracks' table.
*/
const createArtistTrackTableQuery = `
	CREATE TABLE IF NOT EXISTS artist_track  (
	   artist_id UUID REFERENCES artists (id),
	   track_id UUID REFERENCES tracks (id),
	   PRIMARY KEY (artist_id, track_id)
	);
`

// dropArtistTrackTableQuery is a SQL query that drops the 'artist_track' table from the database
const dropArtistTrackTableQuery = `
	DROP TABLE IF EXISTS artist_track;
`

// insertArtistTrackQuery is a SQL query to insert a link from an artist to a track in the 'artist_track' table.
const insertArtistTrackQuery = `
	INSERT INTO artist_track (artist_id, track_id)
    VALUES ($1, $2) 
`

/*
+------------------+-------------------+---------------+
|  artist_id       |  track_id         |    tab_id     |
+------------------+-------------------+---------------+
| 234e5678-e90c... | 111e5251-e90c...  | 123e4567-e... |
| 234e5649-e90c... | 111e5252-e90c...  | 123e4569-e... |
+------------------+-------------------+---------------+

This `v_artists_tracks_tabs` view is a listing of Artist, Track, and Tab information in our system.
The same artist or track may occur multiple times, but a tab is unique.

The view combines these tables and includes the following columns:
  - 'artist_id': The ID of the artist from the `artists` table.
  - 'track_id': The ID of the track from the `tracks` table.
  - 'tab_id': The ID of the tab from the `tabs` table.
*/
const createArtistsTracksTabsViewQuery = `
	CREATE VIEW v_artists_tracks_tabs AS
	SELECT 
		a.id AS artist_id,
		tr.id AS track_id,
		t.id AS tab_id
	FROM artist_track AS at
	INNER JOIN artists AS a ON at.artist_id = a.id
	INNER JOIN track_tab AS tt ON at.track_id = tt.track_id
	INNER JOIN tracks AS tr ON tt.track_id = tr.id
	INNER JOIN tabs AS t ON tt.tab_id = t.id;
`

// dropArtistsTracksTabsViewQuery is a SQL query to drop the 'artists' 'tracks' 'tabs' view if it exists
const dropArtistsTracksTabsViewQuery = `
	DROP VIEW IF EXISTS "v_artists_tracks_tabs";
`

// selectArtistsTracksTabsViewQuery is a SQL query string used to retrieve the artists, tracks and tabs, given a set of IDs, from the database.
const selectArtistsTracksTabsViewQuery = `
	SELECT * FROM v_artists_tracks_tabs WHERE artist_id = ANY($1);
`
