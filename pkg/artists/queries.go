package artists

// insertArtistQuery is a SQL query string used to insert an artist into the 'artists' table.
const insertArtistQuery = `
	INSERT INTO artists (id, name)
    VALUES ($1, $2) 
`

// getArtistsFromIDs is a SQL query string to retrieve the artists with the provided IDs from the 'artists' table.
const getArtistsFromIDs = `SELECT id, name FROM artists WHERE id = ANY($1::uuid[])`
