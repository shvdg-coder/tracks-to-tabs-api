package queries

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

/*
CreateArtistTrackTableQuery is a query to create an 'artist to track' linking table.
+--------------------------------------+--------------------------------------+
|              artist_id               |              track_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

It contains the following columns:
- 'artist_id': The UUID that uniquely identifies an artist in the 'artists' table.
- 'track_id': The UUID that uniquely identifies a track in the 'tracks' table.
*/
const CreateArtistTrackTableQuery = `
	CREATE TABLE IF NOT EXISTS artist_track  (
	   artist_id UUID REFERENCES artists (id),
	   track_id UUID REFERENCES tracks (id),
	   PRIMARY KEY (artist_id, track_id)
	);
`

/*
CreateTracksTableQuery is a query to create the tracks table.
+--------------------------------------+--------------------+------------+
|                 id                   |       title        |  duration  |
+--------------------------------------+--------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |  356000    |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |  483000    |
+--------------------------------------+--------------------+------------+

It contains the following columns:
  - 'id': The UUID that uniquely identifies a track in our system.
  - 'title': The title of the track.
  - 'duration': The duration of the track in milliseconds.
*/
const CreateTracksTableQuery = `
	CREATE TABLE IF NOT EXISTS tracks (
	   id UUID PRIMARY KEY,
	   title VARCHAR(500) NOT NULL,
	   duration NUMERIC NOT NULL
	);
`

/*
CreateTrackTabTableQuery is a query to create the 'track to tab' table.
+--------------------------------------+--------------------------------------+
|               track_id               |                tab_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

- 'track_id': The UUID that uniquely identifies a track from the 'tracks' table.
- 'tab_id': The UUID that uniquely identifies a tab from the 'tabs' table.
*/
const CreateTrackTabTableQuery = `
	CREATE TABLE IF NOT EXISTS track_tab  (
	   track_id UUID REFERENCES tracks (id),
	   tab_id UUID REFERENCES tabs (id),
	   PRIMARY KEY (track_id, tab_id)
	);
`

/*
CreateTabsTableQuery is a query to create a tracks table.
+--------------------------------------+---------------+---------------+------------------+
|                  id                  | instrument_id | difficulty_id |   description    |
+--------------------------------------+---------------+---------------+------------------+
| 123e4567-e89b-12d3-a456-426614174000 |      580      |      423      | James Hetfield   |
| 123e4567-e89b-12d3-a456-426614174001 |      590      |      420      | Mick Mars        |
+--------------------------------------+---------------+---------------+------------------+

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record.
  - 'instrument_id': This column represents the ID of the instrument from a lookup table.
  - 'difficulty_id': This column represents the ID of the difficulty level from a lookup table.
  - 'description': This column records the description of the tab.
*/
const CreateTabsTableQuery = `
	CREATE TABLE IF NOT EXISTS tabs (
	   id UUID PRIMARY KEY,
	   instrument_id INT NOT NULL,
	   difficulty_id INT NOT NULL,
	   description TEXT,
	   CONSTRAINT fk_instrument FOREIGN KEY(instrument_id) REFERENCES instruments(id),
	   CONSTRAINT fk_difficulty	FOREIGN KEY(difficulty_id) REFERENCES difficulties(id)
	);
`

/*
CreateInstrumentsTableQuery is a query to create an instruments lookup table.
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  distortion guitar  |
| 2  |  acoustic guitar    |
+----+---------------------+

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the instrument.
*/
const CreateInstrumentsTableQuery = `
	CREATE TABLE IF NOT EXISTS instruments (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

/*
CreateDifficultiesTableQuery is a query to create a difficulties lookup table.
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  easy               |
| 2  |  intermediate       |
| 3  |  hard               |
+----+---------------------+

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the difficulty level.
*/
const CreateDifficultiesTableQuery = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

/*
CreateReferencesTableQuery is a query to create a references table.
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
const CreateReferencesTableQuery = `
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

/*
CreateSourcesTableQuery is a query to create sources table.
It is used to store external sources.
+-------+-------------------+-------------+
| id    |       name        |  category   |
+-------+-------------------+-------------+
| 1001  | Music Provider 1  |   music     |
| 1002  | Music Provider 2  |   music     |
| 1003  | Tab Provider 1    |   tabs      |
| 1004  | Tab Provider 2    |   tabs      |
+-------+-------------------+-------------+

It contains the following columns:
  - 'id': An auto-incrementing integer that uniquely identifies a record.
  - 'name': The name of the source.
  - 'category': The category of the source.
*/
const CreateSourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS sources(
	   id int PRIMARY KEY,
	   name VARCHAR(250) NOT NULL,
	   category VARCHAR(100) NOT NULL,
	   UNIQUE(name)                                    
	);
`

/*
CreateEndpointsTableQuery is a query to create the endpoints table.
It is used to store endpoints, taken from external sources.
+---------------------------------------------------------------+
|   source_id  | category   | type      | url                   |
+---------------------------------------------------------------+
| 1001         | artist     | web       | /artist/{artistID}    |
| 1001         | track      | web       | /track/{trackID}      |
| 1003         | tab        | api       | /tab/{trackID}        |
+---------------------------------------------------------------+

It contains the following columns:
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type.
  - 'url': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const CreateEndpointsTableQuery = `
	CREATE TABLE IF NOT EXISTS "endpoints" (
	   source_id INT NOT NULL,
	   category VARCHAR(250) NOT NULL,
	   type VARCHAR(250) NOT NULL,
	   url VARCHAR(250) NOT NULL,
	   UNIQUE(source_id, category, type, URL),  
	   CONSTRAINT fk_source FOREIGN KEY(source_id) REFERENCES sources(id)
	);
`

/*
CreateUsersTableQuery is a query to create a users table.
+--------------------------------------+---------------+-------------+
|                   id                 |     email     | password 	 |
+--------------------------------------+---------------+-------------+
| 123e4567-e89b-12d3-a456-426614174000 | john@doe.com  | hashedPw123 |
| 123e4567-e89b-12d3-a456-426614174001 | jane@doe.com  | hashedPw456 |
+--------------------------------------+---------------+-------------+

It consists of the following columns:
  - 'id': This is the UUID that uniquely identifies a user in our system.
  - 'email': This is the user's email address.
  - 'password': This stores the hashed password of the user.
*/
const CreateUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   id UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);
`

/*
CreateSessionsTableQuery is a query to create the sessions table.
It follows the implementation of the postgresstore, from the 'scs' library (https://github.com/alexedwards/scs/tree/master/postgresstore).
It provides Postgres-based storage for HTTP sessions.
+-----------------+----------------------------------------------+---------------------+
|      Token      |                     Data                     |        Expiry       |
+-----------------+----------------------------------------------+---------------------+
| abcdef123456    | 0x537061636573686970... (Byte array data)    | 2023-03-25 09:30:00 |
| ghijkl789012    | 0x4465636f646572736c... (Byte array data)    | 2023-04-11 14:45:00 |
+-----------------+----------------------------------------------+---------------------+

It consists of the following columns:
  - 'Token': A unique session token that serves as the primary key for each session. It matches the sessionID in the scs library.
  - 'Data': Session data stored as binary data (BYTEA). This information corresponds to the encoded and signed session data.
  - 'Expiry': This is the timestamp at which the session is set to expire. It matches the expiry time of the session as handled by the scs library.
*/
const CreateSessionsTableQuery = `
	CREATE TABLE IF NOT EXISTS sessions  (
		token TEXT PRIMARY KEY,
		data BYTEA NOT NULL,
		expiry TIMESTAMPTZ NOT NULL
	);
	`

// CreateSessionExpiryIndexQuery is a SQL query to create an index on the 'sessions' table in the 'expiry' column, if it does not already exist.
const CreateSessionExpiryIndexQuery = `
	DO $$ BEGIN
		IF NOT EXISTS (
			SELECT 1
			FROM   pg_class c
			JOIN   pg_namespace n ON n.oid = c.relnamespace
			WHERE  c.relname = 'sessions_expiry_idx'
			AND    n.nspname = 'public' 
		) THEN
			CREATE INDEX sessions_expiry_idx ON sessions (expiry);
		END IF;
	END $$`
