package tracktab

/*
+--------------------------------------+--------------------------------------+
|               track_id               |                tab_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is used to link the 'tracks' table and the 'tabs' table.

- 'track_id': The UUID that uniquely identifies a track from the 'tracks' table.
- 'tab_id': The UUID that uniquely identifies a tab from the 'tabs' table.
*/
const createTrackTabTableQuery = `
	CREATE TABLE IF NOT EXISTS track_tab  (
	   track_id UUID REFERENCES tracks (id),
	   tab_id UUID REFERENCES tabs (id),
	   PRIMARY KEY (track_id, tab_id)
	);
`

// dropTrackTabTableQuery is a SQL query that drops the 'track_tab' table from the database.
const dropTrackTabTableQuery = `
	DROP TABLE IF EXISTS track_tab;
`

// insertTrackTabQuery is a SQL query to insert a link from a track to a tab in the 'track_tab' table.
const insertTrackTabQuery = `
	INSERT INTO track_tab (track_id, tab_id)
    VALUES ($1, $2) 
`

// getTabIDs retrieves the Tab IDs for the provided Track IDs.
const getTabIDs = `SELECT tab_id FROM track_tab WHERE track_id = ($1)`
