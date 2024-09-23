package dataAPI

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"testing"
)

// TestGetArtists tests whether artists can be inserted and retrieved cascading.
func TestGetArtists(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	svcManager := pkg.NewSvcManager(dbEnv)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	insertCSVFiles(t, dbEnv)

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
	api := pkg.NewDataAPI(svcManager)

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
}
