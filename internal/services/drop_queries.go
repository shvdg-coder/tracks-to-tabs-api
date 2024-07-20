package services

// dropArtistsTableQuery is a SQL query to drop the 'artists' table from the database.
const dropArtistsTableQuery = `
	DROP TABLE IF EXISTS artists;
`

// dropTracksTableQuery is a SQL query to drop the 'tracks' table from the database.
const dropTracksTableQuery = `
	DROP TABLE IF EXISTS tracks;
`

// dropDifficultiesTableQuery is a SQL query to drop the 'difficulties' table if it exists
const dropDifficultiesTableQuery = `
	DROP TABLE IF EXISTS difficulties;
`

// dropEndpointsTableQuery is a SQL query to drop the 'endpoints' table from the database.
const dropEndpointsTableQuery = `
	DROP TABLE IF EXISTS "endpoints";
`

// dropInstrumentsTableQuery is a SQL query to drop the 'instruments' table if it exists
const dropInstrumentsTableQuery = `
	DROP TABLE IF EXISTS instruments;
`

// dropReferencesTableQuery is a SQL query to drop the 'references' table from the database.
const dropReferencesTableQuery = `
	DROP TABLE IF EXISTS "references";
`

// dropSessionsTableQuery is a SQL query to drop the 'sessions' table.
const dropSessionsTableQuery = `
	DROP TABLE IF EXISTS sessions;
`

// dropSourcesTableQuery is a SQL query to drop the 'sources' table
const dropSourcesTableQuery = `
	DROP TABLE IF EXISTS sources;
`

// dropTabsTableQuery is a SQL query to drop the 'tabs' table if it exists
const dropTabsTableQuery = `
	DROP TABLE IF EXISTS tabs;
`

// dropTrackTabTableQuery is a SQL query that drops the 'track_tab' table from the database.
const dropTrackTabTableQuery = `
	DROP TABLE IF EXISTS track_tab;
`

// dropUsersTableQuery is an SQL query to drop the 'users' table from the database.
const dropUsersTableQuery = `
	DROP TABLE IF EXISTS users;
`

// dropArtistTrackTableQuery is a SQL query that drops the 'artist_track' table from the database
const dropArtistTrackTableQuery = `
	DROP TABLE IF EXISTS artist_track;
`
