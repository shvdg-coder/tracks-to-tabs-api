package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	tstenv "github.com/shvdg-dev/tunes-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"testing"
)

// TestArtistsCascading tests whether artists can be inserted and retrieved cascading.
func TestArtistsCascading(t *testing.T) {
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

	// Tests
	artists, err := api.GetArtistsCascading(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist cascading: %s", err.Error())
	}

	if len(artists) != len(artistIDs) {
		t.Fatalf("expected number of artists found in the database (%d) to be equal to those in the CSV (%d)", len(artists), len(artistIDs))
	}
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
