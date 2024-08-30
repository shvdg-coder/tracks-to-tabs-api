package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedTrack contains the data of what a models.Track is expected to have.
type ExpectedTrack struct {
	ID string
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

	testFieldsOfTracks(t, actualTracks, createExpectedTracks(t))
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

// createExpectedTracks constructs and returns a slice of ExpectedTrack objects for use in test cases.
func createExpectedTracks(t *testing.T) []*ExpectedTrack {
	expectedTracks := make([]*ExpectedTrack, 0)

	tracksMap := createTracksFromCSV(t, tracksCSV)
	for id, track := range tracksMap {
		expectedTrack := &ExpectedTrack{
			ID:             id,
			TrackEntry:     track,
			TabCount:       1,
			ReferenceCount: 1,
			ResourceCount:  2,
		}
		expectedTracks = append(expectedTracks, expectedTrack)
	}

	return expectedTracks
}

// createTracksFromCSV creates a map of tracks where the key is the ID and the value a models.TrackEntry.
func createTracksFromCSV(t *testing.T, filePath string) map[string]*models.TrackEntry {
	tracksMap := make(map[string]*models.TrackEntry)

	records, err := logic.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of tracks from a CSV: %s", err.Error())
	}

	for _, record := range records {
		trackID, _ := logic.StringToUUID(record[0])

		track := &models.TrackEntry{
			ID:    trackID,
			Title: record[1],
		}

		tracksMap[trackID.String()] = track
	}

	return tracksMap
}
