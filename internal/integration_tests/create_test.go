package integration_tests

import (
	tstenv "github.com/shvdg-dev/tracks-to-tabs-api/internal/integration_tests/environments"
	"testing"
)

// TestEnvSetup tests whether the tables can be created as part of the database environment setup.
func TestEnvSetup(t *testing.T) {
	dbEnv, err := tstenv.NewService().CreatePostgresEnv()
	if err != nil {
		t.Fatal(err)
	}
	defer dbEnv.Breakdown()

	dbEnv.Setup()
}
