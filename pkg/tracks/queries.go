package tracks

// insertTrackQuery is a SQL query to insert a track into the 'tracks' table.
const insertTrackQuery = `
	INSERT INTO tracks (id, title, duration)
    VALUES ($1, $2, $3) 
`

// getTracksFromIDs is a SQL query to retrieve the tracks with the provided track IDs from the 'tracks' table.
const getTracksFromIDs = `SELECT id, title, duration FROM tracks WHERE id IN ($1)`
