package integration_tests

import (
	"github.com/google/uuid"
	"github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"strconv"
	"testing"
)

// createArtistsFromCSV creates a map of artists where the key is the IDStr and the value a models.ArtistEntry.
func createArtistsFromCSV(t *testing.T, filePath string) map[uuid.UUID]*models.ArtistEntry {
	artistsMap := make(map[uuid.UUID]*models.ArtistEntry)

	records, err := pkg.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of artists from a CSV: %s", err.Error())
	}

	for _, record := range records {
		artistID, _ := pkg.StringToUUID(record[0])

		artist := &models.ArtistEntry{
			ID:   artistID,
			Name: record[1],
		}

		artistsMap[artistID] = artist
	}

	return artistsMap
}

// createTracksFromCSV creates a map of tracks where the key is the IDStr and the value a models.TrackEntry.
func createTracksFromCSV(t *testing.T, filePath string) map[uuid.UUID]*models.TrackEntry {
	tracksMap := make(map[uuid.UUID]*models.TrackEntry)

	records, err := pkg.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of tracks from a CSV: %s", err.Error())
	}

	for _, record := range records {
		trackID, _ := pkg.StringToUUID(record[0])

		track := &models.TrackEntry{
			ID:    trackID,
			Title: record[1],
		}

		tracksMap[trackID] = track
	}

	return tracksMap
}

// createTabsFromCSV creates a map of tabs where the key is the ID and the value a models.TabEntry.
func createTabsFromCSV(t *testing.T, filePath string) map[uuid.UUID]*models.TabEntry {
	tabsMap := make(map[uuid.UUID]*models.TabEntry)

	records, err := pkg.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of tabs from a CSV: %s", err.Error())
	}

	for _, record := range records {
		tabID, _ := pkg.StringToUUID(record[0])

		instrumentID, _ := strconv.ParseUint(record[1], 10, 64)
		difficultyID, _ := strconv.ParseUint(record[2], 10, 64)

		tab := &models.TabEntry{
			ID:           tabID,
			InstrumentID: uint(instrumentID),
			DifficultyID: uint(difficultyID),
			Description:  record[3],
		}

		tabsMap[tabID] = tab
	}

	return tabsMap
}
