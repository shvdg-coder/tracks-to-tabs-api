package tracks

/*
+--------------------------------------+--------------------+
|                 ID                   |       Title        |
+--------------------------------------+--------------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |
+--------------------------------------+--------------------+

This table is used to store the tracks of songs.

It contains the following columns:
  - 'ID': This is the UUID that uniquely identifies a track in our system.
  - 'Title': This column records the Title of the track.
*/
const createTracksTableQuery = `
	CREATE TABLE IF NOT EXISTS tracks (
	   ID UUID PRIMARY KEY,
	   Title VARCHAR(500) NOT NULL
	);
`

// dropTracksTableQuery is a SQL query to drop the 'tracks' table from the database.
const dropTracksTableQuery = `
	DROP TABLE IF EXISTS tracks;
`

// insertTrackQuery is a SQL query to insert a track into the 'tracks' table.
const insertTrackQuery = `
	INSERT INTO tracks (id, title)
    VALUES (gen_random_uuid(), $2) 
`
