package integration_tests

import (
	"encoding/json"
	logic "github.com/shvdg-dev/base-logic/pkg"
	tstenv "github.com/shvdg-dev/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"os"
	"reflect"
	"testing"
)

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := setup(t)
	defer dbEnv.Breakdown()

	// Prepare
	expectedArtistsJson, err := os.ReadFile(artists1JSON)
	if err != nil {
		t.Fatalf("failed to read '%s': %s", artists1JSON, err)
	}

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

	// Convert artists to JSON format
	artistsJSON, err := json.Marshal(artists)
	if err != nil {
		t.Fatalf("error occurred during marshalling to JSON: %s", err.Error())
		return
	}

	// Tests
	isEqual, err := isEqualJSON(string(artistsJSON), string(expectedArtistsJson))
	if err != nil {
		t.Fatalf("failed to compare JSONs: %s", err)
	}

	if !isEqual {
		t.Fatalf("JSONs are not equal: got \n%s", string(artistsJSON))
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

// isEqualJSON checks whether the two provided JSON strings are equal.
func isEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}
	var err error

	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, err
	}

	return reflect.DeepEqual(o1, o2), nil
}
