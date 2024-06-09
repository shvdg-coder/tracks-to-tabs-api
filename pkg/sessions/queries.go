package sessions

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

// dropSessionsTableQuery is a SQL query to drop the 'sessions' table.
const dropSessionsTableQuery = `
	DROP TABLE IF EXISTS sessions;
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
