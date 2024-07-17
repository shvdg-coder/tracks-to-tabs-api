package artists

/*
+--------------------------------------+------------+
|                   id                 |   name    |
+--------------------------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Artist 1   |
| 123e4567-e89b-12d3-a456-426614174001 | Artist 2   |
+--------------------------------------+------------+

This table is used to store Artists in our system.

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record in our system.
  - 'name': This column has the name of an Artist.
*/
const createArtistsTableQuery = `
	CREATE TABLE IF NOT EXISTS artists  (
	   id UUID PRIMARY KEY,
	   name VARCHAR(500) NOT NULL
	);
`

// dropArtistsTableQuery is a SQL query to drop the 'artists' table from the database.
const dropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`

// insertArtistQuery is a SQL query string used to insert an artist into the 'artists' table.
const insertArtistQuery = `
	INSERT INTO artists (id, name)
    VALUES ($1, $2) 
`

// getArtistsFromIDs is a SQL query string to retrieve the artists with the provided IDs from the 'artists' table.
const getArtistsFromIDs = `SELECT id, name FROM artists WHERE id IN ($1)`
