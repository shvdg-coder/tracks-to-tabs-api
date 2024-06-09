package artists

const createArtistsTableQuery = `
	CREATE TABLE IF NOT EXISTS artists  (
	   ID UUID PRIMARY KEY,
	   Name VARCHAR(500) NOT NULL
	);
`

const dropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`

const insertArtistQuery = `
	INSERT INTO artists (id, name)
    VALUES (gen_random_uuid(), $1) 
    ON CONFLICT DO NOTHING;
`
