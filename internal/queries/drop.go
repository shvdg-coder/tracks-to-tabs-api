package queries

// DropArtistsTableQuery is a SQL query to drop the 'artists' table from the database.
const DropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`

// DropArtistTrackTableQuery is a SQL query that drops the 'artist_track' table from the database
const DropArtistTrackTableQuery = `
	DROP TABLE IF EXISTS artist_track;
`

// DropTracksTableQuery is a SQL query to drop the 'tracks' table from the database.
const DropTracksTableQuery = `
	DROP TABLE IF EXISTS tracks;
`

// DropTrackTabTableQuery is a SQL query that drops the 'track_tab' table from the database.
const DropTrackTabTableQuery = `
	DROP TABLE IF EXISTS track_tab;
`

// DropTabsTableQuery is a SQL query to drop the 'tabs' table if it exists
const DropTabsTableQuery = `
	DROP TABLE IF EXISTS tabs;
`

// DropInstrumentsTableQuery is a SQL query to drop the 'instruments' table if it exists
const DropInstrumentsTableQuery = `
	DROP TABLE IF EXISTS instruments;
`

// DropDifficultiesTableQuery is a SQL query to drop the 'difficulties' table if it exists
const DropDifficultiesTableQuery = `
	DROP TABLE IF EXISTS difficulties;
`

// DropReferencesTableQuery is a SQL query to drop the 'references' table from the database.
const DropReferencesTableQuery = `
	DROP TABLE IF EXISTS "references";
`

// DropSourcesTableQuery is a SQL query to drop the 'sources' table
const DropSourcesTableQuery = `
	DROP TABLE IF EXISTS sources;
`

// DropEndpointsTableQuery is a SQL query to drop the 'endpoints' table from the database.
const DropEndpointsTableQuery = `
	DROP TABLE IF EXISTS "endpoints";
`

// DropUsersTableQuery is an SQL query to drop the 'users' table from the database.
const DropUsersTableQuery = `
	DROP TABLE IF EXISTS users;
`

// DropSessionsTableQuery is a SQL query to drop the 'sessions' table.
const DropSessionsTableQuery = `
	DROP TABLE IF EXISTS sessions;
`
