package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedArtist contains the data of what a models.Artist is expected to have.
type ExpectedArtist struct {
	ID             string
	Name           string
	TrackCount     int
	ReferenceCount int
	ResourceCount  int
}

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defaultInsertions(t, dbEnv)
	defer dbEnv.Breakdown()

	// Prepare
	artistIDStrings, err := logic.GetCSVColumnValues(artistsCSV, artistsColumnID)
	if err != nil {
		t.Fatal(err)
	}

	artistIDs, err := logic.StringsToUUIDs(artistIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	api := pkg.NewDataAPI(dbEnv)

	// Execute
	actualArtists, err := api.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist: %s", err.Error())
	}

	// Test
	if len(actualArtists) != len(artistIDs) {
		t.Errorf("expected %d artists, got %d", len(artistIDs), len(actualArtists))
	}

	testFieldsOfArtists(t, actualArtists, createExpectedArtists())
}

// createExpectedArtists constructs and returns a slice of ExpectedArtist objects for use in test cases.
func createExpectedArtists() []*ExpectedArtist {
	return []*ExpectedArtist{
		{ID: "9b5ed980-340a-4b69-8d47-92e358992657", Name: "Knocked Loose", TrackCount: 2, ReferenceCount: 3, ResourceCount: 2},
		{ID: "1d7edc80-526b-4f74-a5de-657a7c9e0521", Name: "Ice Nine Kills", TrackCount: 2, ReferenceCount: 3, ResourceCount: 2},
	}
}

// testFieldsOfArtists tests the fields of multiple artist objects by comparing the actual artists to the expected ones.
func testFieldsOfArtists(t *testing.T, actualArtists []*models.Artist, expectedArtist []*ExpectedArtist) {
	for i := 0; i < len(actualArtists); i++ {
		testFieldsOfArtist(t, actualArtists[i], expectedArtist[i])
	}
}

// testFieldsOfArtist tests the fields of a single artist object by comparing the actual artist to the expected one.
func testFieldsOfArtist(t *testing.T, actualArtist *models.Artist, expectedArtist *ExpectedArtist) {
	// Check ID
	artistID, _ := logic.UUIDToString(actualArtist.ID)
	if artistID != expectedArtist.ID {
		t.Errorf("expected ID to be %s, got %s", expectedArtist.ID, actualArtist.ID)
	}

	// Check Artist Title
	if actualArtist.Name != expectedArtist.Name {
		t.Errorf("expected Title to be %s, got %s", expectedArtist.Name, actualArtist.Name)
	}

	// Check Tracks
	if len(actualArtist.Tracks) != expectedArtist.TrackCount {
		t.Errorf("expected %d Tracks, got %d", expectedArtist.TrackCount, len(actualArtist.Tracks))
	}

	// Check References
	if len(actualArtist.References) != expectedArtist.ReferenceCount {
		t.Errorf("expected %d References, got %d", expectedArtist.ReferenceCount, len(actualArtist.References))
	}

	// Check Resources
	if len(actualArtist.Resources) != expectedArtist.ResourceCount {
		t.Errorf("expected %d Resources, got %d", expectedArtist.ResourceCount, len(actualArtist.Resources))
	}
}
