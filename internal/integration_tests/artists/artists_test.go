package integration_tests

import (
	tstenv "github.com/shvdg-dev/tunes-to-tabs-api/internal/integration_tests/environments"
	"log"
	"testing"
)

// TestArtistsCascading tests whether artists can be inserted and retrieved cascading.
func TestArtistsCascading(t *testing.T) {
	// Setup
	dbEnv, err := tstenv.NewService().CreatePostgresEnv()
	if err != nil {
		t.Fatal(err)
	}
	defer dbEnv.Breakdown()
	dbEnv.Setup()

	// Inserting data
	err = dbEnv.InsertCSVFile(artistsCSV, artistsTable, artistsFields)
	if err != nil {
		log.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(tracksCSV, tracksTable, tracksFields)
	if err != nil {
		log.Fatal(err)
	}
	err = dbEnv.InsertCSVFile(artisttrackCSV, artisttrackTable, artisttrackFields)
	if err != nil {
		log.Fatal(err)
	}

	// Tests
	// TODO:
}
