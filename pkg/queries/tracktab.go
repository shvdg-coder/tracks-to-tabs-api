package queries

// InsertTrackTab is a SQL query to insert a link from a track to a tab in the 'track_tab' table.
const InsertTrackTab = `
	INSERT INTO track_tab (track_id, tab_id)
    VALUES ($1, $2) 
`

// GetTrackTabLinks retrieves the 'track to tab' links for the provided Track IDs.
const GetTrackTabLinks = `SELECT track_id, tab_id FROM track_tab WHERE track_id = ANY($1::uuid[])`
