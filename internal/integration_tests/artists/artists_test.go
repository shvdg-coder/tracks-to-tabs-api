package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	tstenv "github.com/shvdg-dev/tunes-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	trk "github.com/shvdg-dev/tunes-to-tabs-api/pkg/models"
	"testing"
)

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := setup(t)
	defer dbEnv.Breakdown()

	// Prepare
	artistIDStrings, err := logic.GetCSVColumnValues(artistsCSV, artistsColumnID)
	if err != nil {
		t.Fatal(err)
	}
	artistIDs, err := logic.StringsToUUIDs(artistIDStrings...)
	if err != nil {
		t.Fatal(err)
	}
	api := pkg.NewAPI(dbEnv)

	// Execute
	artists, err := api.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist cascading: %s", err.Error())
	}

	// Extract
	tracks := extractTracks(artists)
	tabs := extractTabs(tracks)

	// Tests
	testArtists(t, artists)
	testTracks(t, tracks)
	testTabs(t, tabs)
}

// testArtists todo:
func testArtists(t *testing.T, artists []*trk.Artist) {
	if len(artists) != 2 {
		t.Fatalf("expected number of artists found in the database (%d) to be equal to those in the CSV (%d)", len(artists), 2)
	}
}

// testTracks todo:
func testTracks(t *testing.T, tracks []*trk.Track) {
	if len(tracks) != 4 {
		t.Fatalf("expected number of tracks found in the database (%d) to be equal to those in the CSV (%d)", len(tracks), 4)
	}
}

// testTabs todo:
func testTabs(t *testing.T, tabs []*trk.Tab) {
	if len(tabs) != 4 {
		t.Fatalf("expected number of tabs found in the database (%d) to be equal to those in the CSV (%d)", len(tabs), 4)
	}
}

// extractTracks extracts the tracks.Track's from the artists.ArtistEntry.
func extractTracks(artists []*trk.Artist) []*trk.Track {
	var tracks []*trk.Track
	for _, artist := range artists {
		tracks = append(tracks, artist.Tracks...)
	}
	return tracks
}

// extractTabs extracts the tabs.Tab's from the tracks.Track.
func extractTabs(tracks []*trk.Track) []*trk.Tab {
	var tabs []*trk.Tab
	for _, track := range tracks {
		tabs = append(tabs, track.Tabs...)
	}
	return tabs
}

// setup prepares the tests by performing the minimally required steps.
func setup(t *testing.T) tstenv.DbEnvOperations {
	dbEnv, err := tstenv.NewService().CreatePostgresEnv()
	if err != nil {
		t.Fatal(err)
	}
	dbEnv.Setup()
	insertions(t, dbEnv)
	return dbEnv
}

// insertions prepares the test, by insertions the dummy data into the database.
func insertions(t *testing.T, dbEnv tstenv.DbEnvOperations) {
	err := dbEnv.InsertCSVFile(artistsCSV, artistsTable, artistsColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(tracksCSV, tracksTable, tracksColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(artisttrackCSV, artistTrackTable, artisttrackColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(sourcesCSV, sourcesTable, sourcesColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(endpointsCSV, endpointsTable, endpointsColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(instrumentsCSV, instrumentsTable, instrumentsColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(difficultiesCSV, difficultiesTable, difficultiesColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(tabsCSV, tabsTable, tabsColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(tracktabCSV, tabTrackTable, trackTabColumns)
	if err != nil {
		t.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(referencesCSV, referencesTable, referencesColumns)
	if err != nil {
		t.Fatal(err)
	}
}
