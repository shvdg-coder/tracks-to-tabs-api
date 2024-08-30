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
	   PRIMARY KEY (track_id, tab_id)
	);
`

// DropTrackTabTable is a SQL query that drops the 'track_tab' table from the database.
const DropTrackTabTable = `
	DROP TABLE IF EXISTS track_tab;
`

// InsertTrackTab is a SQL query to insert a link from a track to a tab in the 'track_tab' table.
const InsertTrackTab = `
	INSERT INTO track_tab (track_id, tab_id)
    VALUES ($1, $2) 
`

// GetTrackTabLinks retrieves the 'track to tab' links for the provided Track IDs.
const GetTrackTabLinks = `SELECT track_id, tab_id FROM track_tab WHERE track_id = ANY($1::uuid[]) OR tab_id = ANY($1::uuid[])`
