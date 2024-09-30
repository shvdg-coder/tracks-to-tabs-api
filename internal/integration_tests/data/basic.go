package dataAPI

import (
	tstenv "github.com/shvdg-coder/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/services"
	"testing"
)

// createDefaultDbEnv prepares the tests by performing the minimally required steps.
func createDefaultDbEnv(t *testing.T) tstenv.DbEnvOps {
	dbEnv, err := tstenv.NewEnvSvc().CreatePostgresEnv()
	if err != nil {
		t.Fatal(err)
	}
	dbEnv.Setup()
	return dbEnv
}

// seed seeds the database using the provided configuration.
func seed(t *testing.T, dbEnv tstenv.DbEnvOps, apiConfigPath string) *models.SeedingConfig {
	apiConfig, err := models.NewAPIConfig(apiConfigPath)
	if err != nil {
		t.Fatalf("error occurred while parsing the seed config: %s", err.Error())
	}

	seedingConfig := apiConfig.Seeding
	svcManager := services.NewSvcManager(dbEnv)

	dummyAPI := services.NewDummySvc(svcManager, seedingConfig.Sources, seedingConfig.Instruments, seedingConfig.Difficulties)
	seedingAPI := services.NewSeedSvc(svcManager, seedingConfig, dummyAPI)
	seedingAPI.Seed()

	return seedingConfig
}

// insertCSVFiles prepares the test, by insertCSVFiles the dummy data into the database.
func insertCSVFiles(t *testing.T, dbEnv tstenv.DbEnvOps) {
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
