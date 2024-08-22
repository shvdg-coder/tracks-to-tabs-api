package queries

// InsertTrack is a SQL query to insert a track into the 'tracks' table.
const InsertTrack = `
	INSERT INTO tracks (id, title, duration)
    VALUES ($1, $2, $3) 
`

// GetTracksFromIDs is a SQL query to retrieve the tracks with the provided track IDs from the 'tracks' table.
const GetTracksFromIDs = `SELECT id, title, duration FROM tracks WHERE id = ANY($1::uuid[])`
