package integration_tests

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedTrack contains the data of what a models.Track is expected to have.
type ExpectedTrack struct {
	*models.TrackEntry
	TabCount       int
	ReferenceCount int
	ResourceCount  int
}

// TestGetTracks tests whether tracks can be inserted and retrieved cascading.
func TestGetTracks(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	defaultData(t, dbEnv)

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
	api := pkg.NewDataAPI(dbEnv)

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
	// Check IDStr
	if actualTrack.ID != expectedTrack.ID {
		t.Errorf("expected IDStr to be %s, got %s", expectedTrack.ID, actualTrack.ID)
	}

	// Check Track Title
	if actualTrack.Title != expectedTrack.Title {
		t.Errorf("expected Title to be %s, got %s", expectedTrack.Title, actualTrack.Title)
	}

	// Check Tabs
	if len(actualTrack.Tabs) != expectedTrack.TabCount {
		t.Errorf("expected %d Tabs, got %d", expectedTrack.TabCount, len(actualTrack.Tabs))
	}

	// Check References
	if len(actualTrack.References) != expectedTrack.ReferenceCount {
		t.Errorf("expected %d References, got %d", expectedTrack.ReferenceCount, len(actualTrack.References))
	}

	// Check Resources
	if len(actualTrack.Resources) != expectedTrack.ResourceCount {
		t.Errorf("expected %d Resources, got %d", expectedTrack.ResourceCount, len(actualTrack.Resources))
	}
}

// createExpectedTracks constructs and returns a map of ExpectedTrack's for use in test cases.
func createExpectedTracks(t *testing.T) map[uuid.UUID]*ExpectedTrack {
	expectedTracks := make(map[uuid.UUID]*ExpectedTrack)

	tracksMap := createTracksFromCSV(t, tracksCSV)
	for id, track := range tracksMap {
		expectedTrack := &ExpectedTrack{
			TrackEntry:     track,
			TabCount:       1,
			ReferenceCount: 1,
			ResourceCount:  2,
		}
		expectedTracks[id] = expectedTrack
	}

	return expectedTracks
}
