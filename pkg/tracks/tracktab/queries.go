package tracktab

// insertTrackTabQuery is a SQL query to insert a link from a track to a tab in the 'track_tab' table.
const insertTrackTabQuery = `
	INSERT INTO track_tab (track_id, tab_id)
    VALUES ($1, $2) 
`

// getTabIDs retrieves the Tab IDs for the provided Track IDs.
const getTabIDs = `SELECT tab_id FROM track_tab WHERE track_id = ($1)`
