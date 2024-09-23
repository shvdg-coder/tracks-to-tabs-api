package queries

/*
CreateArtistTrackTable is a query to create an 'artist to track' linking table.
+--------------------------------------+--------------------------------------+
|              artist_id               |              track_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

It contains the following columns:
- 'artist_id': The UUID that uniquely identifies an artist in the 'artists' table.
- 'track_id': The UUID that uniquely identifies a track in the 'tracks' table.
*/
const CreateArtistTrackTable = `
	CREATE TABLE IF NOT EXISTS artist_track  (
	   artist_id UUID REFERENCES artists(id),
	   track_id UUID REFERENCES tracks(id),
	   PRIMARY KEY (artist_id, track_id),
	   CONSTRAINT fk_artist FOREIGN KEY(artist_id) REFERENCES artists(id),
	   CONSTRAINT fk_track FOREIGN KEY(track_id) REFERENCES tracks(id)
	);

	CREATE INDEX idx_artisttrack_artist_id ON artist_track(artist_id);
	CREATE INDEX idx_artisttrack_track_id ON artist_track(track_id);
`

// DropArtistTrackTableQuery is a SQL query that drops the 'artist_track' table from the database
const DropArtistTrackTableQuery = `
	DROP TABLE IF EXISTS artist_track;
`

// GetArtistTrackLinks is for retrieving 'artist to track' links for the provided IDs.
const GetArtistTrackLinks = `SELECT artist_id, track_id FROM artist_track WHERE artist_id = ANY($1::uuid[]) OR track_id = ANY($1::uuid[])`
