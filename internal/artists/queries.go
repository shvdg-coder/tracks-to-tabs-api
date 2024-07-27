package artists

/*
CreateArtistsTableQuery is a query to create an artists table.
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
const CreateArtistsTableQuery = `
	CREATE TABLE IF NOT EXISTS artists  (
	   id UUID PRIMARY KEY,
	   name VARCHAR(500) NOT NULL
	);
`

// DropArtistsTableQuery is a SQL query to drop the 'artists' table from the database.
const DropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`
