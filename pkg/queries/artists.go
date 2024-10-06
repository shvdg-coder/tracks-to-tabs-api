package queries

/*
CreateArtistsTable is a query to create an artists table.
+--------------------------------------+------------+--------------------------------------+--------------------------------------+
|                   id                 |   name     |                 cover                |                banner                |
+--------------------------------------+------------+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174000 | Artist 1   | cover_image_url_1                    | banner_image_url_1                   |
| 123e4567-e89b-12d3-a456-426614174001 | Artist 2   | cover_image_url_2                    | banner_image_url_2                   |
+--------------------------------------+------------+--------------------------------------+--------------------------------------+

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record in our system.
  - 'name': This column has the name of an Artist.
  - 'cover': This column contains the cover image URL as a string.
  - 'banner': This column contains the banner image URL as a string.
*/
const CreateArtistsTable = `
	CREATE TABLE IF NOT EXISTS artists  (
	   id UUID PRIMARY KEY,
	   name VARCHAR(500) NOT NULL,
	   cover TEXT,
	   banner TEXT
	);

	CREATE INDEX idx_artist_id ON "artists"(id);
`

// DropArtistsTable is a SQL query to drop the 'artists' table from the database.
const DropArtistsTable = `
	DROP TABLE IF EXISTS artists;
`

// GetArtistsFromIDs is a SQL query string to retrieve the artists with the provided IDs from the 'artists' table.
const GetArtistsFromIDs = `SELECT id, name, cover, banner FROM artists WHERE id = ANY($1::uuid[])`
