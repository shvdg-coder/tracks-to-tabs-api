package resources

/*
+---------------------------------------------------------------------------------------------------------+
|           InternalID             |             SourceID    		  | Category   | Type  |  Resource    |
+---------------------------------------------------------------------------------------------------------+
| 123e4567-e89b-12d3-a456-42661417 | 123e4567-e89b-12d3-a456-42661419 | Artist     | ID    | SP123        |
| 123e4567-e89b-12d3-a456-42661418 | 123e4567-e89b-12d3-a456-42661419 | Artist     | Image | someone.jpg  |
| 123e4567-e89b-12d3-a456-42661420 | 123e4567-e89b-12d3-a456-42661419 | Playlist   | ID    | ST123        |
| 123e4567-e89b-12d3-a456-42661421 | 123e4567-e89b-12d3-a456-42661419 | Tab        | ID    | YT123        |
| 123e4567-e89b-12d3-a456-42661422 | 123e4567-e89b-12d3-a456-42661419 | Tab        | ID    | UG123        |
+---------------------------------------------------------------------------------------------------------+

The table 'resources' is used to store references of various resources for internal records.

It contains the following columns:
  - 'InternalID': This is the UUID of a record in our system.
  - 'SourceID': This is the UUID of the external source from which the data was referenced.
  - 'Category': This denotes the category of an external resource (e.g., 'Artist', 'Playlist', 'Tab').
  - 'Type': This denotes the type of the resource (e.g., 'Endpoint', 'Image', 'ID').
  - 'Resource': This stores the actual resource data (e.g. "SP123", "someone.jpg", "ST123", "YT123", "UG123").
*/
const createResourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS resources(
	   InternalID UUID NOT NULL,
	   SourceID UUID NOT NULL,
	   Category VARCHAR(250) NOT NULL,
	   Type VARCHAR(250) NOT NULL,
	   Resource VARCHAR(250) NOT NULL,
	   UNIQUE(InternalID, Category, Type)                                    
	);
`

// dropResourcesTableQuery is a SQL query to drop the 'resources' table from the database.
const dropResourcesTableQuery = `
	DROP TABLE IF EXISTS resources;
`

// insertResourceQuery is a SQL query to insert a new resource record in the 'resources' table
const insertResourceQuery = `
	INSERT INTO resources (InternalID, SourceID, Category, Type, Resource)
    VALUES ($1, $2, $3, $4, $5) 
`
