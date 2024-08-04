package integration_tests

// CSV
const artistsCSV = "./resources/artists.csv"
const tracksCSV = "./resources/tracks.csv"
const artisttrackCSV = "./resources/artisttrack.csv"

// Tables
const artistsTable = "artists"
const tracksTable = "tracks"
const artisttrackTable = "artist_track"

// Columns
const artistsColumnID = "id"
const artistsColumnName = "name"
const tracksColumnID = "id"
const tracksColumnTitle = "title"
const tracksColumnDuration = "duration"
const artistTrackColumnArtistID = "artist_id"
const artistTrackColumnTrackID = "track_id"

var artistsColumns = []string{artistsColumnID, artistsColumnName}
var tracksColumns = []string{tracksColumnID, tracksColumnTitle, tracksColumnDuration}
var artisttrackColumns = []string{artistTrackColumnArtistID, artistTrackColumnTrackID}
