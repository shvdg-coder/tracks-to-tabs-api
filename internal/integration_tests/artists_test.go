package integration_tests

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedArtist contains the data of what a models.Artist is expected to have.
type ExpectedArtist struct {
	*models.ArtistEntry
	TrackCount     int
	ReferenceCount int
	ResourceCount  int
}

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	defaultData(t, dbEnv)

	expectedArtistsMap := createExpectedArtists(t)

	artistIDStrings, err := logic.GetCSVColumnValues(artistsCSV, artistsColumnID)
	if err != nil {
		t.Fatal(err)
	}

	artistIDs, err := logic.StringsToUUIDs(artistIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	mapper := mappers.NewArtistSvc()
	api := pkg.NewDataAPI(dbEnv)

	// Execute
	actualArtists, err := api.GetArtists(artistIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of artist: %s", err.Error())
	}

	// Test
	if len(actualArtists) == 0 || len(artistIDs) == 0 {
		t.Errorf("expected more than 0 artists (CSV: %d, API: %d)", len(artistIDs), len(actualArtists))
	}

	if len(actualArtists) != len(artistIDs) {
		t.Errorf("expected %d artists, got %d", len(artistIDs), len(actualArtists))
	}

	testFieldsOfArtists(t, mapper.ArtistsToMap(actualArtists), expectedArtistsMap)
}

// testFieldsOfArtists tests the fields of multiple artist objects by comparing the actual artists to the expected ones.
func testFieldsOfArtists(t *testing.T, actualArtistsMap map[uuid.UUID]*models.Artist, expectedArtistsMap map[uuid.UUID]*ExpectedArtist) {
	for id := range actualArtistsMap {
		actualArtist := actualArtistsMap[id]
		expectedArtist, ok := expectedArtistsMap[id]
		if !ok {
			t.Fatalf("ID %s does not exist in 'expected artists' map", id)
		} else {
			testFieldsOfArtist(t, actualArtist, expectedArtist)
		}
	}
}

// testFieldsOfArtist tests the fields of a single artist object by comparing the actual artist to the expected one.
func testFieldsOfArtist(t *testing.T, actualArtist *models.Artist, expectedArtist *ExpectedArtist) {
	// Check ID
	if actualArtist.ID != expectedArtist.ID {
		t.Errorf("expected ID to be %s, got %s", expectedArtist.ID, actualArtist.ID)
	}

	// Check Artist Name
	if actualArtist.Name != expectedArtist.Name {
		t.Errorf("expected Name to be %s, got %s", expectedArtist.Name, actualArtist.Name)
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

// createExpectedArtists constructs and returns a map of ExpectedArtist's for use in test cases.
func createExpectedArtists(t *testing.T) map[uuid.UUID]*ExpectedArtist {
	expectedArtists := make(map[uuid.UUID]*ExpectedArtist)

	artistsMap := createArtistsFromCSV(t, artistsCSV)
	for id, artist := range artistsMap {
		expectedArtist := &ExpectedArtist{
			ArtistEntry:    artist,
			TrackCount:     2,
			ReferenceCount: 3,
			ResourceCount:  2,
		}
		expectedArtists[id] = expectedArtist
	}

	return expectedArtists
}

// createArtistsFromCSV creates a map of artists where the key is the IDStr and the value a models.ArtistEntry.
func createArtistsFromCSV(t *testing.T, filePath string) map[uuid.UUID]*models.ArtistEntry {
	artistsMap := make(map[uuid.UUID]*models.ArtistEntry)

	records, err := logic.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of artists from a CSV: %s", err.Error())
	}

	for _, record := range records {
		artistID, _ := logic.StringToUUID(record[0])

		artist := &models.ArtistEntry{
			ID:   artistID,
			Name: record[1],
		}

		artistsMap[artistID] = artist
	}

	return artistsMap
}
