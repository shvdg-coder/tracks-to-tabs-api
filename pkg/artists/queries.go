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
+------------------+--------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
|  artist_id       | artist_name  |  track_id    |   track_title      |  track_duration  |     tab_id    |   instrument_id   | instrument_name  | difficulty_id  | difficulty_name  | description          |
+---------------------------------+--------------+--------------------+------------------+---------------+-------------------+------------------+----------------+------------------+----------------------+
| 234e5678-e90c... | Sting        |     600      | Englishman in NY   |   140305         | 123e4567-e... |  700              |  Bass guitar     | 500            | Intermediate     | Sting playing bass   |
| 234e5678-e90c... | B.B. King    |     650      | The Thrill is Gone |   124022         | 123e4569-e... |  785              |  Electric guitar | 530            | Advanced         | B.B. King's solo     |
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
const createArtistsToTabsViewQuery = `
	CREATE VIEW v_artists_to_tabs AS
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

// dropArtistsToTabsViewQuery is a SQL query to drop the 'artists' to 'tabs' view if it exists
const dropArtistsToTabsViewQuery = `
	DROP VIEW IF EXISTS "v_artists_to_tabs";
`
