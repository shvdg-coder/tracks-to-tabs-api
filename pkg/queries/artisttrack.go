package queries

// InsertArtistTrack is a SQL query to insert a link from an artist to a track in the 'artist_track' table.
const InsertArtistTrack = `
	INSERT INTO artist_track (artist_id, track_id)
    VALUES ($1, $2) 
`

// GetArtistTrackLinks is for retrieving 'artist to track' links for the provided artist IDs.
const GetArtistTrackLinks = `SELECT artist_id, track_id FROM artist_track WHERE artist_id = ANY($1::uuid[])`
