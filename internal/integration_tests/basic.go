package integration_tests

import (
	tstenv "github.com/shvdg-dev/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// createDefaultDbEnv prepares the tests by performing the minimally required steps.
func createDefaultDbEnv(t *testing.T) tstenv.DbEnvOperations {
	dbEnv, err := tstenv.NewService().CreatePostgresEnv()
	if err != nil {
		t.Fatal(err)
	}
	dbEnv.Setup()
	return dbEnv
}

// seed seeds the database using the provided configuration.
func seed(t *testing.T, dbEnv tstenv.DbEnvOperations, seedConfigPath string) {
	seedConfig, err := models.NewSeedConfig(seedConfigPath)
	if err != nil {
		t.Fatalf("error occurred while parsing the seed config: %s", err.Error())
	}

	dummyAPI := pkg.NewDummyAPI(dbEnv, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)
	seedingAPI := pkg.NewSeedingAPI(dbEnv, seedConfig, dummyAPI)
	seedingAPI.Seed()
}

// defaultData prepares the test, by defaultData the dummy data into the database.
func defaultData(t *testing.T, dbEnv tstenv.DbEnvOperations) {
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
