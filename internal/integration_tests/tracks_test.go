package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedTrack contains the data of what a models.Track is expected to have.
type ExpectedTrack struct {
	ID             string
	Title          string
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

	trackIDStrings, err := logic.GetCSVColumnValues(tracksCSV, tracksColumnID)
	if err != nil {
		t.Fatal(err)
	}

	trackIDs, err := logic.StringsToUUIDs(trackIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	api := pkg.NewDataAPI(dbEnv)

	// Execute
	actualTracks, err := api.GetTracks(trackIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of track: %s", err.Error())
	}

	// Test
	if len(actualTracks) != len(trackIDs) {
		t.Errorf("expected %d tracks, got %d", len(trackIDs), len(actualTracks))
	}

	testFieldsOfTracks(t, actualTracks, createExpectedTracks())
}

// createExpectedTracks constructs and returns a slice of ExpectedTrack objects for use in test cases.
func createExpectedTracks() []*ExpectedTrack {
	return []*ExpectedTrack{
		{ID: "c51a9150-6b7d-45aa-88f7-75372b221c1d", Title: "Suffocate", TabCount: 1, ReferenceCount: 1, ResourceCount: 2},
		{ID: "c52a9150-6b7d-45aa-88f7-75372b222c1e", Title: "Blinding Faith", TabCount: 1, ReferenceCount: 1, ResourceCount: 2},
		{ID: "c72a9150-6b7d-45aa-88f7-75372b222f1f", Title: "Stabbing In The Dark", TabCount: 1, ReferenceCount: 1, ResourceCount: 2},
		{ID: "c73a9150-6b7d-45aa-88f7-75372b223f1d", Title: "Ex-Mørtis", TabCount: 1, ReferenceCount: 1, ResourceCount: 2},
	}
}

// testFieldsOfTracks tests the fields of multiple track objects by comparing the actual tracks to the expected ones.
func testFieldsOfTracks(t *testing.T, actualTracks []*models.Track, expectedTrack []*ExpectedTrack) {
	for i := 0; i < len(actualTracks); i++ {
		testFieldsOfTrack(t, actualTracks[i], expectedTrack[i])
	}
}

// testFieldsOfTrack tests the fields of a single track object by comparing the actual track to the expected one.
func testFieldsOfTrack(t *testing.T, actualTrack *models.Track, expectedTrack *ExpectedTrack) {
	// Check ID
	trackID, _ := logic.UUIDToString(actualTrack.ID)
	if trackID != expectedTrack.ID {
		t.Errorf("expected ID to be %s, got %s", expectedTrack.ID, actualTrack.ID)
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
