package queries

/*
CreateArtistsTable is a query to create an artists table.
+--------------------------------------+------------+
|                   id                 |   name    |
+--------------------------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Artist 1   |
| 123e4567-e89b-12d3-a456-426614174001 | Artist 2   |
+--------------------------------------+------------+

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record in our system.
  - 'name': This column has the name of an Artist.
*/
const CreateArtistsTable = `
	CREATE TABLE IF NOT EXISTS artists  (
	   id UUID PRIMARY KEY,
	   name VARCHAR(500) NOT NULL
	);
`

// DropArtistsTable is a SQL query to drop the 'artists' table from the database.
const DropArtistsTable = `
	DROP TABLE IF EXISTS artists;
`

// InsertArtists is a SQL query string used to insert artist(s) into the 'artists' table.
const InsertArtists = `
	INSERT INTO artists (id, name)
    VALUES ($1, $2) 
`

// GetArtistsFromIDs is a SQL query string to retrieve the artists with the provided IDs from the 'artists' table.
const GetArtistsFromIDs = `SELECT id, name FROM artists WHERE id = ANY($1::uuid[])`
