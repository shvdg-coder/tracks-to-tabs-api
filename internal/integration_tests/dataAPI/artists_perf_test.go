package dataAPI

import (
	"github.com/google/uuid"
	env "github.com/shvdg-coder/tracks-to-tabs-api/internal/integration_tests/environments"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg"
	"testing"
	"time"
)

// TestGetArtistsPerf tests the performance of retrieving artists.
func TestGetArtistsPerf(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	svcManager := pkg.NewSvcManager(dbEnv)
	defer dbEnv.Breakdown()

	// Prepare
	start := time.Now()
	seedConfig := seed(t, dbEnv, maxConfigPath)
	elapsed := time.Since(start)
	t.Logf("Seeding artists took %s", elapsed.Round(time.Millisecond))

	api := pkg.NewDataAPI(svcManager)

	// Execute
	artistIDs := selectArtistIDs(t, dbEnv)

	start = time.Now()
	artists, err := api.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist: %s", err.Error())
	}

	elapsed = time.Since(start)
	artistsCount, tracksCount, tabsCount := countRecords(t, dbEnv)
	totalRecordCount := artistsCount + tracksCount + tabsCount
	t.Logf("GetArtists took %s (total: %d, artists: %d, tracks: %d, tabs: %d)",
		elapsed.Round(time.Millisecond),
		totalRecordCount, artistsCount, tracksCount, tabsCount)

	// Test
	if len(artists) == 0 || len(artistIDs) == 0 {
		t.Errorf("expected more than 0 artists (DB: %d, API: %d)", len(artistIDs), len(artists))
	}

	if len(artists) < seedConfig.Dummies.Artists.Min {
		t.Errorf("expected a minimum of %d artists, got %d", seedConfig.Dummies.Artists.Min, len(artists))
	}
}

// selectArtistIDs retrieves al the artist IDs from the artists table.
func selectArtistIDs(t *testing.T, dbEnv env.DbEnvOperations) []uuid.UUID {
	rows, err := dbEnv.DB().Query("SELECT id FROM artists")
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

// countRecords count the records of the artists, tracks and tabs.
func countRecords(t *testing.T, dbEnv env.DbEnvOperations) (artistCount int, trackCount int, tabCount int) {
	err := dbEnv.DB().QueryRow("SELECT COUNT(*) FROM artists").Scan(&artistCount)
	if err != nil {
		t.Fatalf("Failed to count artists: %s", err.Error())
	}

	err = dbEnv.DB().QueryRow("SELECT COUNT(*) FROM tracks").Scan(&trackCount)
	if err != nil {
		t.Fatalf("Failed to count tracks: %s", err.Error())
	}

	err = dbEnv.DB().QueryRow("SELECT COUNT(*) FROM tabs").Scan(&tabCount)
	if err != nil {
		t.Fatalf("Failed to count tabs: %s", err.Error())
	}

	return artistCount, trackCount, tabCount
}
