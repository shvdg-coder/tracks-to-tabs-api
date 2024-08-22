package queries

// InsertArtist is a SQL query string used to insert an artist into the 'artists' table.
const InsertArtist = `
	INSERT INTO artists (id, name)
    VALUES ($1, $2) 
`

// GetArtistsFromIDs is a SQL query string to retrieve the artists with the provided IDs from the 'artists' table.
const GetArtistsFromIDs = `SELECT id, name FROM artists WHERE id = ANY($1::uuid[])`
