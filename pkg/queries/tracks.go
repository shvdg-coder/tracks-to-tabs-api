package queries

/*
CreateTracksTableQuery is a query to create the tracks table.
+--------------------------------------+--------------------+------------+
|                 id                   |       title        |  duration  |
+--------------------------------------+--------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |  356000    |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |  483000    |
+--------------------------------------+--------------------+------------+

It contains the following columns:
  - 'id': The UUID that uniquely identifies a track in our system.
  - 'title': The title of the track.
  - 'duration': The duration of the track in milliseconds.
*/
const CreateTracksTableQuery = `
	CREATE TABLE IF NOT EXISTS tracks (
	   id UUID PRIMARY KEY,
	   title VARCHAR(500) NOT NULL,
	   duration NUMERIC NOT NULL
	);
`

// DropTracksTableQuery is a SQL query to drop the 'tracks' table from the database.
const DropTracksTableQuery = `
	DROP TABLE IF EXISTS tracks;
`

// InsertTrack is a SQL query to insert a track into the 'tracks' table.
const InsertTrack = `
	INSERT INTO tracks (id, title, duration)
    VALUES ($1, $2, $3) 
`

// GetTracksFromIDs is a SQL query to retrieve the tracks with the provided track IDs from the 'tracks' table.
const GetTracksFromIDs = `SELECT id, title, duration FROM tracks WHERE id = ANY($1::uuid[])`
