package integration_tests

// CSV
const artistsCSV = "./resources/artists.csv"
const tracksCSV = "./resources/tracks.csv"
const artisttrackCSV = "./resources/artisttrack.csv"
const tabsCSV = "./resources/tabs.csv"
const tracktabCSV = "./resources/tracktab.csv"
const endpointsCSV = "./resources/endpoints.csv"
const referencesCSV = "./resources/references.csv"
const sourcesCSV = "./resources/sources.csv"
const instrumentsCSV = "./resources/instruments.csv"
const difficultiesCSV = "./resources/difficulties.csv"

// Expected outputs
const artists1JSON = "./expectations/artists_1.json"

// Tables
const artistsTable = "artists"
const tracksTable = "tracks"
const artistTrackTable = "artist_track"
const tabsTable = "tabs"
const tabTrackTable = "track_tab"
const endpointsTable = "endpoints"
const referencesTable = "references"
const sourcesTable = "sources"
const instrumentsTable = "instruments"
const difficultiesTable = "difficulties"

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

// Columns: Instruments
const instrumentsColumnID = "id"
const instrumentsColumnName = "name"

// Columns: Difficulties
const difficultiesColumnID = "id"
const difficultiesColumnName = "name"

// Columns: Tabs
const tabsColumnID = "id"
const tabsColumnInstrumentID = "instrument_id"
const tabsColumnDifficultyID = "difficulty_id"
const tabsColumnDescription = "description"

// Columns: Track Tabs
const trackTabColumnTrackID = "track_id"
const trackTabColumnTabID = "tab_id"

var artistsColumns = []string{artistsColumnID, artistsColumnName}
var tracksColumns = []string{tracksColumnID, tracksColumnTitle, tracksColumnDuration}
var artisttrackColumns = []string{artistTrackColumnArtistID, artistTrackColumnTrackID}
var endpointsColumns = []string{endpointsColumnSourceID, endpointsColumnCategory, endpointsColumnType, endpointsColumnURL}
var referencesColumns = []string{referencesColumnInternalID, referencesColumnSourceID, referencesColumnCategory, referencesColumnType, referencesColumnReference}
var sourcesColumns = []string{sourcesColumnSourceID, sourcesColumnSourceName, sourcesColumnSourceCategory}
var instrumentsColumns = []string{instrumentsColumnID, instrumentsColumnName}
var difficultiesColumns = []string{difficultiesColumnID, difficultiesColumnName}
var tabsColumns = []string{tabsColumnID, tabsColumnInstrumentID, tabsColumnDifficultyID, tabsColumnDescription}
var trackTabColumns = []string{trackTabColumnTrackID, trackTabColumnTabID}
