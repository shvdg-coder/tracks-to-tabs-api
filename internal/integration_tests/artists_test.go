package integration_tests

import (
	"encoding/json"
	logic "github.com/shvdg-dev/base-logic/pkg"
	tstenv "github.com/shvdg-dev/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"os"
	"reflect"
	"testing"
)

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := setup(t)
	defer dbEnv.Breakdown()

	// Prepare
	expectedArtists := getExpectedArtists(t)

	artistIDStrings, err := logic.GetCSVColumnValues(artistsCSV, artistsColumnID)
	if err != nil {
		t.Fatal(err)
	}

	artistIDs, err := logic.StringsToUUIDs(artistIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	api := pkg.NewDataAPI(dbEnv)

	// Execute
	artists, err := api.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist cascading: %s", err.Error())
	}

	// Tests
	if len(artists) != len(expectedArtists) {
		t.Fatalf("expected to be the same number of artists")
	}

	for i := range artists {
		if !reflect.DeepEqual(artists[i], expectedArtists[i]) {
			actualJSON, _ := json.Marshal(artists[i])
			expectedJSON, _ := json.Marshal(expectedArtists[i])
			t.Errorf("expected artist \n%s, \nbut got \n%s", string(expectedJSON), string(actualJSON))
		}
	}
}

// getExpectedArtists unmarshalls the expected artists from the JSON.
func getExpectedArtists(t *testing.T) []models.Artist {
	expectedArtistsJson, err := os.ReadFile(artistsScenario1JSON)
	if err != nil {
		t.Fatalf("failed to read '%s': %s", artistsScenario1JSON, err)
	}

	// Unmarshal expected artists from JSON
	var expectedArtists []models.Artist
	if err := json.Unmarshal(expectedArtistsJson, &expectedArtists); err != nil {
		t.Fatalf("failed to unmarshal expected artists: %s", err)
	}

	return expectedArtists
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
