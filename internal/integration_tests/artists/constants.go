package integration_tests

// CSV
const artistsCSV = "./artists.csv"
const tracksCSV = "./tracks.csv"
const artisttrackCSV = "./artisttrack.csv"

// Tables
const artistsTable = "artists"
const tracksTable = "tracks"
const artisttrackTable = "artist_track"

// Fields
var artistsFields = []string{"id", "name"}
var tracksFields = []string{"id", "title", "duration"}
var artisttrackFields = []string{"artist_id", "track_id"}
