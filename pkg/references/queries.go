package references

/*
+-------------------------------------------------------------------------------------+
|           internal_id            |   source_id  | category   | type  | reference    |
+-------------------------------------------------------------------------------------+
| 123e4567-e89b-12d3-a456-42661417 | 1001         | Artist     | ID    | SP123        |
| 123e4567-e89b-12d3-a456-42661418 | 1001         | Artist     | Image | someone.jpg  |
| 123e4567-e89b-12d3-a456-42661420 | 1002         | Track      | ID    | YT123        |
| 123e4567-e89b-12d3-a456-42661421 | 1003         | Tab        | ID    | ST123        |
| 123e4567-e89b-12d3-a456-42661422 | 1003         | Tab        | ID    | UG123        |
+-------------------------------------------------------------------------------------+

The table 'references' is used to store references of various references for internal records.

It contains the following columns:
  - 'internal_id': This is the UUID of a record in our system.
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference (e.g., 'Artist', 'Track', 'Tab').
  - 'type': This denotes the type of the reference (e.g., 'Image', 'ID').
  - 'reference': This stores the actual reference data (e.g. "SP123", "someone.jpg", "ST123", "YT123", "UG123").
*/
const createReferencesTableQuery = `
	CREATE TABLE IF NOT EXISTS "references" (
	   internal_id UUID NOT NULL,
	   source_id INT NOT NULL,
	   category VARCHAR(250) NOT NULL, 
	   type VARCHAR(250) NOT NULL,
	   reference VARCHAR(250) NOT NULL,
	   UNIQUE(internal_id, source_id, category, type),
       CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
	);
`

// dropReferencesTableQuery is a SQL query to drop the 'references' table from the database.
const dropReferencesTableQuery = `
	DROP TABLE IF EXISTS "references";
`

// insertReferenceQuery is a SQL query to insert a new reference record in the 'references' table
const insertReferenceQuery = `
	INSERT INTO "references" (internal_id, source_id, category, type, reference)
    VALUES ($1, $2, $3, $4, $5) 
`
