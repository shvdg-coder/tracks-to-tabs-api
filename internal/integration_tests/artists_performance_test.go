package integration_tests

import (
	"github.com/google/uuid"
	env "github.com/shvdg-dev/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
	"time"
)

// TestGetArtistsPerformance tests the performance of retrieving artists.
func TestGetArtistsPerformance(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defer dbEnv.Breakdown()

	// Prepare
	seedConfig, err := models.NewSeedConfig(performanceConfigPath)
	if err != nil {
		t.Fatalf("error occurred while parsing the seed config: %s", err.Error())
	}

	dummyAPI := pkg.NewDummyAPI(dbEnv, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)
	seedingAPI := pkg.NewSeedingAPI(dbEnv, seedConfig, dummyAPI)
	dataAPI := pkg.NewDataAPI(dbEnv)

	// Execute
	seedingAPI.Seed()
	artistIDs := selectArtistIDs(t, dbEnv)

	start := time.Now()
	artists, err := dataAPI.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist: %s", err.Error())
	}

	elapsed := time.Since(start)
	t.Logf("GetArtists took %s", elapsed.Round(time.Millisecond))

	// Test
	if len(artists) != len(artistIDs) {
		t.Errorf("expected %d number of artists, got %d", len(artistIDs), len(artists))
	}
}

// selectArtistIDs retrieves al the artist IDs from the artists table.
func selectArtistIDs(t *testing.T, dbEnv env.DbEnvOperations) []uuid.UUID {
	rows, err := dbEnv.Query("SELECT id FROM artists")
	if err != nil {
		t.Fatalf("error occured while querying artists table: %s", err.Error())
	}

	artistIDs := make([]uuid.UUID, 0)
	for rows.Next() {
		artistID := uuid.UUID{}
		err := rows.Scan(&artistID)
		if err != nil {
			t.Fatalf("error occured while scanning artists rows: %s", err.Error())
		}
		artistIDs = append(artistIDs, artistID)
	}

	return artistIDs
}
