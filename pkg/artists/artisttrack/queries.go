package artisttrack

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

// getArtistTrackFromIDs is for retrieving 'artist to track' links for the provided artist IDs.
const getArtistTrackFromIDs = `SELECT artist_id, track_id FROM artist_track WHERE artist_id IN ($1)`
