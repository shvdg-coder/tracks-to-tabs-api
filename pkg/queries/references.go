package queries

/*
CreateReferencesTable is a query to create a references table.
It is used to store references of various external sources to link them to internal records.
+-------------------------------------------------------------------------------------+
|           internal_id            |   source_id  | category   | type  | reference    |
+-------------------------------------------------------------------------------------+
| 123e4567-e89b-12d3-a456-42661417 | 1001         | artist     | id    | SP123        |
| 123e4567-e89b-12d3-a456-42661418 | 1001         | artist     | image | someone.jpg  |
| 123e4567-e89b-12d3-a456-42661420 | 1002         | track      | id    | YT123        |
| 123e4567-e89b-12d3-a456-42661421 | 1003         | tab        | id    | ST123        |
| 123e4567-e89b-12d3-a456-42661422 | 1003         | tab        | id    | UG123        |
+-------------------------------------------------------------------------------------+

It contains the following columns:
  - 'internal_id': This is the UUID of a record in our system.
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type of the reference.
  - 'reference': This stores the actual reference data.
*/
const CreateReferencesTable = `
	CREATE TABLE IF NOT EXISTS "references" (
	   internal_id UUID NOT NULL,
	   source_id INT REFERENCES sources(id) NOT NULL,
	   category VARCHAR(250) NOT NULL, 
	   type VARCHAR(250) NOT NULL,
	   reference VARCHAR(250) NOT NULL,
	   PRIMARY KEY (internal_id, source_id, category, type)
	);

	CREATE INDEX idx_internal_id ON "references"(internal_id);
`

// DropReferencesTable is a SQL query to drop the 'references' table from the database.
const DropReferencesTable = `
	DROP TABLE IF EXISTS "references";
`

// GetReferences is a SQL query to retrieve references from the provided internal IDs.
const GetReferences = `SELECT internal_id, source_id, category, type, reference FROM "references" WHERE internal_id = ANY($1::uuid[])`
