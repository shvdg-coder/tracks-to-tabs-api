package references

/*
+-------------------------------------------------------------------------------------+
|           InternalID             |   SourceID   | Category   | Type  | Reference    |
+-------------------------------------------------------------------------------------+
| 123e4567-e89b-12d3-a456-42661417 | 1001         | Artist     | ID    | SP123        |
| 123e4567-e89b-12d3-a456-42661418 | 1001         | Artist     | Image | someone.jpg  |
| 123e4567-e89b-12d3-a456-42661420 | 1002         | Track      | ID    | YT123        |
| 123e4567-e89b-12d3-a456-42661421 | 1003         | Tab        | ID    | ST123        |
| 123e4567-e89b-12d3-a456-42661422 | 1003         | Tab        | ID    | UG123        |
+-------------------------------------------------------------------------------------+

The table 'references' is used to store references of various references for internal records.

It contains the following columns:
  - 'InternalID': This is the UUID of a record in our system.
  - 'SourceID': This is the ID of the external source from which the data was referenced.
  - 'Category': This denotes the category of an external reference (e.g., 'Artist', 'Track', 'Tab').
  - 'Type': This denotes the type of the reference (e.g., 'Image', 'ID').
  - 'Reference': This stores the actual reference data (e.g. "SP123", "someone.jpg", "ST123", "YT123", "UG123").
*/
const createReferencesTableQuery = `
	CREATE TABLE IF NOT EXISTS "references" (
	   InternalID UUID NOT NULL,
	   SourceID INT NOT NULL,
	   Category VARCHAR(250) NOT NULL, 
	   Type VARCHAR(250) NOT NULL,
	   Reference VARCHAR(250) NOT NULL,
	   UNIQUE(InternalID, SourceID, Category, Type),
       CONSTRAINT fk_source FOREIGN KEY(SourceID) REFERENCES sources(ID)
	);
`

// dropReferencesTableQuery is a SQL query to drop the 'references' table from the database.
const dropReferencesTableQuery = `
	DROP TABLE IF EXISTS "references";
`

// insertReferenceQuery is a SQL query to insert a new reference record in the 'references' table
const insertReferenceQuery = `
	INSERT INTO "references" (InternalID, SourceID, Category, Type, Reference)
    VALUES ($1, $2, $3, $4, $5) 
`
