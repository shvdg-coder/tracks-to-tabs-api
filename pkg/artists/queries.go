package artists

/*
+--------------------------------------+------------+
|                   ID                 |    Name    |
+--------------------------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Artist 1   |
| 123e4567-e89b-12d3-a456-426614174001 | Artist 2   |
+--------------------------------------+------------+

This table is used to store Artists in our system.

It contains the following columns:
  - 'ID': This is the UUID that uniquely identifies a record in our system.
  - 'Name': This column has the name of an Artist.
*/
const createArtistsTableQuery = `
	CREATE TABLE IF NOT EXISTS artists  (
	   ID UUID PRIMARY KEY,
	   Name VARCHAR(500) NOT NULL
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
|              ArtistID                |               TrackID                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is linking the 'artists' table and the 'tracks' table.

- 'ArtistID': The UUID that uniquely identifies an artist in the 'artists' table.
- 'TrackID': The UUID that uniquely identifies a track in the 'tracks' table.
*/
const createArtistTrackTableQuery = `
	CREATE TABLE IF NOT EXISTS artist_track  (
	   ArtistID UUID REFERENCES artists (ID),
	   TrackID UUID REFERENCES tracks (ID),
	   PRIMARY KEY (ArtistID, TrackID)
	);
`

// dropArtistTrackTableQuery is a SQL query that drops the 'artist_track' table from the database
const dropArtistTrackTableQuery = `
	DROP TABLE IF EXISTS artist_track;
`

// insertArtistTrackQuery is a SQL query to insert a link from an artist to a track in the 'artist_track' table.
const insertArtistTrackQuery = `
	INSERT INTO artist_track (artistId, trackId)
    VALUES ($1, $2) 
`
