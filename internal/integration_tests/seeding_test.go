package integration_tests

import (
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// TestSeeding tests whether a database can be seeded from the seed config.
func TestSeeding(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defer dbEnv.Breakdown()

	// Prepare
	seedConfig, err := models.NewSeedConfig(seedConfigPath)
	if err != nil {
		t.Fatalf("error occurred while parsing the seed config: %s", err.Error())
	}

	dummyAPI := pkg.NewDummyAPI(dbEnv, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)
	seedingAPI := pkg.NewSeedingAPI(dbEnv, seedConfig, dummyAPI)

	// Execute
	seedingAPI.Seed()
}
