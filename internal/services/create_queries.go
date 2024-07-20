package services

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

/*
+--------------------------------------+--------------------+------------+
|                 id                   |       title        |  duration  |
+--------------------------------------+--------------------+------------+
| 123e4567-e89b-12d3-a456-426614174000 | Sweet Child O'Mine |  356000    |
| 123e4567-e89b-12d3-a456-426614174001 | Stairway to Heaven |  483000    |
+--------------------------------------+--------------------+------------+

This table is used to store the tracks of songs.

It contains the following columns:
  - 'id': The UUID that uniquely identifies a track in our system.
  - 'title': The title of the track.
  - 'duration': The duration of the track in milliseconds.
*/
const createTracksTableQuery = `
	CREATE TABLE IF NOT EXISTS tracks (
	   id UUID PRIMARY KEY,
	   title VARCHAR(500) NOT NULL,
	   duration NUMERIC NOT NULL
	);
`

/*
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  easy               |
| 2  |  intermediate       |
| 3  |  hard               |
+----+---------------------+

This table is used to store Difficulty Levels in our system.

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the difficulty level.
*/
const createDifficultiesTableQuery = `
	CREATE TABLE IF NOT EXISTS difficulties (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

/*
+---------------------------------------------------------------+
|   source_id  | category   | type      | url                   |
+---------------------------------------------------------------+
| 1001         | artist     | web       | /artist/{artistID}    |
| 1001         | track      | web       | /track/{trackID}      |
| 1003         | tab        | api       | /tab/{trackID}        |
+---------------------------------------------------------------+

The table 'endpoints' is used to store various endpoints for internal records.

It contains the following columns:
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type.
  - 'url': This is the endpoint, which has to be formatted with the corresponding IDs/references, as stored in the 'references' table.
*/
const createEndpointsTableQuery = `
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
+----+---------------------+
| id |         name        |
+----+---------------------+
| 1  |  distortion guitar  |
| 2  |  acoustic guitar    |
+----+---------------------+

This table is used to store Instruments in our system.

It contains the following columns:
  - 'id': This is an auto-incrementing integer that uniquely identifies a record.
  - 'name': This column records the name of the instrument.
*/
const createInstrumentsTableQuery = `
	CREATE TABLE IF NOT EXISTS instruments (
	   id SERIAL PRIMARY KEY,
	   name VARCHAR(255) NOT NULL
	);
`

/*
+-------------------------------------------------------------------------------------+
|           internal_id            |   source_id  | category   | type  | reference    |
+-------------------------------------------------------------------------------------+
| 123e4567-e89b-12d3-a456-42661417 | 1001         | artist     | id    | SP123        |
| 123e4567-e89b-12d3-a456-42661418 | 1001         | artist     | image | someone.jpg  |
| 123e4567-e89b-12d3-a456-42661420 | 1002         | track      | id    | YT123        |
| 123e4567-e89b-12d3-a456-42661421 | 1003         | tab        | id    | ST123        |
| 123e4567-e89b-12d3-a456-42661422 | 1003         | tab        | id    | UG123        |
+-------------------------------------------------------------------------------------+

The table 'references' is used to store references of various references for internal records.

It contains the following columns:
  - 'internal_id': This is the UUID of a record in our system.
  - 'source_id': This is the ID of the external source from which the data was referenced.
  - 'category': This denotes the category of an external reference.
  - 'type': This denotes the type of the reference.
  - 'reference': This stores the actual reference data.
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

/*
+-----------------+----------------------------------------------+---------------------+
|      Token      |                     Data                     |        Expiry       |
+-----------------+----------------------------------------------+---------------------+
| abcdef123456    | 0x537061636573686970... (Byte array data)    | 2023-03-25 09:30:00 |
| ghijkl789012    | 0x4465636f646572736c... (Byte array data)    | 2023-04-11 14:45:00 |
+-----------------+----------------------------------------------+---------------------+

This table is used to store session details in our system.
It follows the implementation of the postgresstore from the 'scs' library (https://github.com/alexedwards/scs/tree/master/postgresstore).
It provides Postgres-based storage for HTTP sessions.

It consists of the following columns:
  - 'Token': A unique session token that serves as the primary key for each session. It matches the sessionID in the scs library.
  - 'Data': Session data stored as binary data (BYTEA). This information corresponds to the encoded and signed session data.
  - 'Expiry': This is the timestamp at which the session is set to expire. It matches the expiry time of the session as handled by the scs library.
*/
const createSessionsTableQuery = `
	CREATE TABLE IF NOT EXISTS sessions  (
		token TEXT PRIMARY KEY,
		data BYTEA NOT NULL,
		expiry TIMESTAMPTZ NOT NULL
	);
	`

// createSessionExpiryIndexQuery is a SQL query to create an index on the 'sessions' table in the 'expiry' column, if it does not already exist.
const createSessionExpiryIndexQuery = `
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

/*
+-------+-------------------+-------------+
| id    |       name        |  category   |
+-------+-------------------+-------------+
| 1001  | Music Provider 1  |   music     |
| 1002  | Music Provider 2  |   music     |
| 1003  | Tab Provider 1    |   tabs      |
| 1004  | Tab Provider 2    |   tabs      |
+-------+-------------------+-------------+

The table named 'sources' has the purpose of storing unique source names.

It contains the following columns:
  - 'id': An auto-incrementing integer that uniquely identifies a record.
  - 'name': The name of the source.
  - 'category': The category of the source.
*/
const createSourcesTableQuery = `
	CREATE TABLE IF NOT EXISTS sources(
	   id int PRIMARY KEY,
	   name VARCHAR(250) NOT NULL,
	   category VARCHAR(100) NOT NULL,
	   UNIQUE(name)                                    
	);
`

/*
+--------------------------------------+---------------+---------------+------------------+
|                  id                  | instrument_id | difficulty_id |   description    |
+--------------------------------------+---------------+---------------+------------------+
| 123e4567-e89b-12d3-a456-426614174000 |      580      |      423      | James Hetfield   |
| 123e4567-e89b-12d3-a456-426614174001 |      590      |      420      | Mick Mars        |
+--------------------------------------+---------------+---------------+------------------+

This table is used to store Tracks in our system.

It contains the following columns:
  - 'id': This is the UUID that uniquely identifies a record.
  - 'instrument_id': This column represents the ID of the instrument from a lookup table.
  - 'difficulty_id': This column represents the ID of the difficulty level from a lookup table.
  - 'description': This column records the description of the tab.
*/
const createTabsTableQuery = `
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
+--------------------------------------+--------------------------------------+
|               track_id               |                tab_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is used to link the 'tracks' table and the 'tabs' table.

- 'track_id': The UUID that uniquely identifies a track from the 'tracks' table.
- 'tab_id': The UUID that uniquely identifies a tab from the 'tabs' table.
*/
const createTrackTabTableQuery = `
	CREATE TABLE IF NOT EXISTS track_tab  (
	   track_id UUID REFERENCES tracks (id),
	   tab_id UUID REFERENCES tabs (id),
	   PRIMARY KEY (track_id, tab_id)
	);
`

/*
+--------------------------------------+---------------+-------------+
|                   id                 |     email     | password 	 |
+--------------------------------------+---------------+-------------+
| 123e4567-e89b-12d3-a456-426614174000 | john@doe.com  | hashedPw123 |
| 123e4567-e89b-12d3-a456-426614174001 | jane@doe.com  | hashedPw456 |
+--------------------------------------+---------------+-------------+

This table is used to store a user with their credentials in our system.

It consists of the following columns:
  - 'id': This is the UUID that uniquely identifies a user in our system.
  - 'email': This is the user's email address.
  - 'password': This stores the hashed password of the user.
*/
const createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users  (
	   id UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(60) NOT NULL
	);
`

/*
+--------------------------------------+--------------------------------------+
|              artist_id               |              track_id                |
+--------------------------------------+--------------------------------------+
| 123e4567-e89b-12d3-a456-426614174050 | 123e4567-e89b-12d3-a456-426614174000 |
| 123e4567-e89b-12d3-a456-426614174051 | 123e4567-e89b-12d3-a456-426614174001 |
+--------------------------------------+--------------------------------------+

This table is linking the 'artists' table and the 'tracks' table.

- 'artist_id': The UUID that uniquely identifies an artist in the 'artists' table.
- 'track_id': The UUID that uniquely identifies a track in the 'tracks' table.
*/
const createArtistTrackTableQuery = `
	CREATE TABLE IF NOT EXISTS artist_track  (
	   artist_id UUID REFERENCES artists (id),
	   track_id UUID REFERENCES tracks (id),
	   PRIMARY KEY (artist_id, track_id)
	);
`
