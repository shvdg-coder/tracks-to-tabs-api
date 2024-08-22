package artisttrack

// insertArtistTrackQuery is a SQL query to insert a link from an artist to a track in the 'artist_track' table.
const insertArtistTrackQuery = `
	INSERT INTO artist_track (artist_id, track_id)
    VALUES ($1, $2) 
`

// getArtistTrackLinks is for retrieving 'artist to track' links for the provided artist IDs.
const getArtistTrackLinks = `SELECT artist_id, track_id FROM artist_track WHERE artist_id = ANY($1::uuid[])`
