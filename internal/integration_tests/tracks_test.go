package integration_tests

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"testing"
)

// TestGetTracks tests whether tracks can be inserted and retrieved cascading.
func TestGetTracks(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	svcManager := pkg.NewSvcManager(dbEnv)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	insertCSVFiles(t, dbEnv)

	expectedTracks := createExpectedTracks(t)

	trackIDStrings, err := logic.GetCSVColumnValues(tracksCSV, tracksColumnID)
	if err != nil {
		t.Fatal(err)
	}

	trackIDs, err := logic.StringsToUUIDs(trackIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	mapper := mappers.NewTrackSvc()
	api := pkg.NewDataAPI(svcManager)

	// Execute
	actualTracks, err := api.GetTracks(trackIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of track: %s", err.Error())
	}

	// Test
	if len(actualTracks) == 0 || len(trackIDs) == 0 {
		t.Errorf("expected more than 0 tracks (CSV: %d, API: %d)", len(trackIDs), len(actualTracks))
	}

	if len(actualTracks) != len(trackIDs) {
		t.Errorf("expected %d tracks, got %d", len(trackIDs), len(actualTracks))
	}

	testFieldsOfTracks(t, mapper.TracksToMap(actualTracks), expectedTracks)
}

// testFieldsOfTracks tests the fields of multiple track objects by comparing the actual tracks to the expected ones.
func testFieldsOfTracks(t *testing.T, actualTracksMap map[uuid.UUID]*models.Track, expectedTracksMap map[uuid.UUID]*ExpectedTrack) {
	for id := range actualTracksMap {
		actualTrack := actualTracksMap[id]
		expectedTrack, ok := expectedTracksMap[id]
		if !ok {
			t.Fatalf("ID %s does not exist in 'expected track' map", id)
		} else {
			testFieldsOfTrack(t, actualTrack, expectedTrack)
		}
	}
}

// testFieldsOfTrack tests the fields of a single track object by comparing the actual track to the expected one.
func testFieldsOfTrack(t *testing.T, actualTrack *models.Track, expectedTrack *ExpectedTrack) {
	// Check ID
	if actualTrack.ID != expectedTrack.ID {
		t.Errorf("expected IDStr to be %s, got %s", expectedTrack.ID, actualTrack.ID)
	}

	// Check Track Title
	if actualTrack.Title != expectedTrack.Title {
		t.Errorf("expected Title to be %s, got %s", expectedTrack.Title, actualTrack.Title)
	}
}
