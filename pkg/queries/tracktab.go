package queries

/*
CreateTrackTabTable is a query to create the 'track to tab' table.
+--------------------------------------+--------------------------------------+
|               track_id               |                tab_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

- 'track_id': The UUID that uniquely identifies a track from the 'tracks' table.
- 'tab_id': The UUID that uniquely identifies a tab from the 'tabs' table.
*/
const CreateTrackTabTable = `
	CREATE TABLE IF NOT EXISTS track_tab  (
	   track_id UUID REFERENCES tracks (id),
	   tab_id UUID REFERENCES tabs (id),
	   PRIMARY KEY (track_id, tab_id),
	   CONSTRAINT fk_track FOREIGN KEY(track_id) REFERENCES tracks(id),
	   CONSTRAINT fk_tab FOREIGN KEY(tab_id) REFERENCES tabs(id)
	);

	CREATE INDEX idx_tracktab_track_id ON track_tab(track_id);
	CREATE INDEX idx_tracktab_tab_id ON track_tab(tab_id);
`

// DropTrackTabTable is a SQL query that drops the 'track_tab' table from the database.
const DropTrackTabTable = `
	DROP TABLE IF EXISTS track_tab;
`

// GetTrackTabLinks retrieves the 'track to tab' links for the provided Track IDs.
const GetTrackTabLinks = `SELECT track_id, tab_id FROM track_tab WHERE track_id = ANY($1::uuid[]) OR tab_id = ANY($1::uuid[])`
