package id_references

/*
+-----------------+--------------------------------------+------------------+------------+
| InternalSource  |          InternalID                  | ExternalSource   | ExternalID |
+-----------------+--------------------------------------+------------------+------------+
| Artist          | 123e4567-e89b-12d3-a456-426614174000 | Spotify          | SP123      |
| Artist          | 123e4567-e89b-12d3-a456-426614174001 | Apple Music      | AM123      |
| Tab             | 123e4567-e89b-12d3-a456-426614174002 | Songster         | SS123      |
| Tab             | 123e4567-e89b-12d3-a456-426614174003 | Ultimate Guitar  | UG123      |
+-----------------+--------------------------------------+------------------+------------+

This table is used to store references of IDs from internal resources and external sources.

It contains the following columns:
  - 'InternalSource': This is the type of the internal object (e.g., 'Artist', 'Tab').
  - 'InternalID': This is the UUID of a record in our system.
  - 'ExternalSource': This is the medium from which the data was referenced, for e.g., Spotify, Apple Music, Songster, Ultimate Guitar etc.
  - 'ExternalID': The unique ID specific to the 'ExternalSource' that identifies the particular data.
*/
const createIdReferencesTableQuery = `
	CREATE TABLE IF NOT EXISTS id_references(
	   InternalSource VARCHAR(250) NOT NULL,
	   InternalID UUID NOT NULL,
	   ExternalSource VARCHAR(250) NOT NULL,
	   ExternalID VARCHAR(250) NOT NULL,
	   UNIQUE(InternalID, ExternalID)                                    
	);
`

// dropIdReferencesTableQuery is a SQL query to drop the 'id_references' table from the database.
const dropIdReferencesTableQuery = `
	DROP TABLE IF EXISTS id_references;
`

// insertIdReferenceQuery is a SQL query to insert a new external reference in the 'id_reference' table
const insertIdReferenceQuery = `
	INSERT INTO id_references (internalSource, internalID, externalsource, externalID)
    VALUES ($1, $2, $3, $4) 
`
