package integration_tests

// CSV
const artistsCSV = "./resources/artists.csv"
const tracksCSV = "./resources/tracks.csv"
const artisttrackCSV = "./resources/artisttrack.csv"
const endpointsCSV = "./resources/endpoints.csv"
const referencesCSV = "./resources/references.csv"
const sourcesCSV = "./resources/sources.csv"

// Tables
const artistsTable = "artists"
const tracksTable = "tracks"
const artisttrackTable = "artist_track"
const endpointsTable = "endpoints"
const referencesTable = "references"
const sourcesTable = "sources"

// Columns: Artists
const artistsColumnID = "id"
const artistsColumnName = "name"

// Columns: Tracks
const tracksColumnID = "id"
const tracksColumnTitle = "title"
const tracksColumnDuration = "duration"

// Columns: Artist Track
const artistTrackColumnArtistID = "artist_id"
const artistTrackColumnTrackID = "track_id"

// Columns: Endpoints
const endpointsColumnSourceID = "source_id"
const endpointsColumnCategory = "category"
const endpointsColumnType = "type"
const endpointsColumnURL = "url"

// Columns: References
const referencesColumnInternalID = "internal_id"
const referencesColumnSourceID = "source_id"
const referencesColumnCategory = "category"
const referencesColumnType = "type"
const referencesColumnReference = "reference"

// Columns: Sources
const sourcesColumnSourceID = "id"
const sourcesColumnSourceName = "name"
const sourcesColumnSourceCategory = "category"

var artistsColumns = []string{artistsColumnID, artistsColumnName}
var tracksColumns = []string{tracksColumnID, tracksColumnTitle, tracksColumnDuration}
var artisttrackColumns = []string{artistTrackColumnArtistID, artistTrackColumnTrackID}
var endpointsColumns = []string{endpointsColumnSourceID, endpointsColumnCategory, endpointsColumnType, endpointsColumnURL}
var referencesColumns = []string{referencesColumnInternalID, referencesColumnSourceID, referencesColumnCategory, referencesColumnType, referencesColumnReference}
var sourcesColumns = []string{sourcesColumnSourceID, sourcesColumnSourceName, sourcesColumnSourceCategory}
