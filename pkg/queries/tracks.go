package queries

/*
CreateTracksTable is a query to create the tracks table.
+--------------------------------------+--------------------+------------+--------------------------------------+
|                 id                   |       title        |  duration  |                 cover                |
+--------------------------------------+--------------------+------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |  356000    | cover_image_url_1                    |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |  483000    | cover_image_url_2                    |
+--------------------------------------+--------------------+------------+--------------------------------------+

It contains the following columns:
  - 'id': The UUID that uniquely identifies a track in our system.
  - 'title': The title of the track.
  - 'duration': The duration of the track in milliseconds.
  - 'cover': This column contains the cover image URL as text.
*/
const CreateTracksTable = `
	CREATE TABLE IF NOT EXISTS tracks (
	   id UUID PRIMARY KEY,
	   title VARCHAR(500) NOT NULL,
	   duration NUMERIC NOT NULL,
	   popularity INT NULL,
	   cover TEXT
	);

	CREATE INDEX idx_track_id ON tracks(id);
`

// DropTracksTable is a SQL query to drop the 'tracks' table from the database.
const DropTracksTable = `
	DROP TABLE IF EXISTS tracks;
`

// GetTracksFromIDs is a SQL query to retrieve the tracks with the provided track IDs from the 'tracks' table.
const GetTracksFromIDs = `SELECT id, title, duration, cover FROM tracks WHERE id = ANY($1::uuid[])`
