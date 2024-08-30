package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"strconv"
	"testing"
)

// ExpectedTab contains the data of what a models.Tab is expected to have.
type ExpectedTab struct {
	ID string
	*models.TabEntry
	ReferencesCount int
	ResourceCount   int
}

// TestGetTabs tests whether tabs can be inserted and retrieved cascading.
func TestGetTabs(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	defaultData(t, dbEnv)

	tabIDStrings, err := logic.GetCSVColumnValues(tabsCSV, tabsColumnID)
	if err != nil {
		t.Fatal(err)
	}

	tabIDs, err := logic.StringsToUUIDs(tabIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	api := pkg.NewDataAPI(dbEnv)

	// Execute
	actualTabs, err := api.GetTabs(tabIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of tab: %s", err.Error())
	}

	// Test
	if len(actualTabs) != len(tabIDs) {
		t.Errorf("expected %d tabs, got %d", len(tabIDs), len(actualTabs))
	}

	testFieldsOfTabs(t, actualTabs, createExpectedTabs(t))
}

// testFieldsOfTabs tests the fields of multiple tab objects by comparing the actual tabs to the expected ones.
func testFieldsOfTabs(t *testing.T, actualTabs []*models.Tab, expectedTab []*ExpectedTab) {
	for i := 0; i < len(actualTabs); i++ {
		testFieldsOfTab(t, actualTabs[i], expectedTab[i])
	}
}

// testFieldsOfTab tests the fields of a single tab object by comparing the actual tab to the expected one.
func testFieldsOfTab(t *testing.T, actualTab *models.Tab, expectedTab *ExpectedTab) {
	// Check ID
	tabID, err := logic.UUIDToString(actualTab.ID)
	if err != nil {
		t.Fatalf("error occurred during conversion of UUID to string: %s", err.Error())
	}

	if tabID != expectedTab.ID {
		t.Errorf("expected ID to be %s, got %s", expectedTab.ID, actualTab.ID)
	}

	// Check InstrumentID
	if actualTab.InstrumentID != expectedTab.InstrumentID {
		t.Errorf("expected InstrumentID to be %d, got %d", expectedTab.InstrumentID, actualTab.InstrumentID)
	}

	// Check DifficultyID
	if actualTab.DifficultyID != expectedTab.DifficultyID {
		t.Errorf("expected DifficultyID to be %d, got %d", expectedTab.DifficultyID, actualTab.DifficultyID)
	}

	// Check ReferencesCount
	if len(actualTab.References) != expectedTab.ReferencesCount {
		t.Errorf("expected %d References, got %d", expectedTab.ReferencesCount, len(actualTab.References))
	}

	// Check ResourceCount
	if len(actualTab.Resources) != expectedTab.ResourceCount {
		t.Errorf("expected %d Resources, got %d", expectedTab.ResourceCount, len(actualTab.Resources))
	}
}

// createExpectedTabs constructs and returns a slice of ExpectedTab objects for use in test cases.
func createExpectedTabs(t *testing.T) []*ExpectedTab {
	expectedTabs := make([]*ExpectedTab, 0)

	tabsMap := createTabsFromCSV(t, tabsCSV)
	for id, tab := range tabsMap {
		expectedTab := &ExpectedTab{
			ID:              id,
			TabEntry:        tab,
			ReferencesCount: 2,
			ResourceCount:   1,
		}
		expectedTabs = append(expectedTabs, expectedTab)
	}

	return expectedTabs
}

// createTabsFromCSV creates a map of tabs where the key is the ID and the value a models.TabEntry.
func createTabsFromCSV(t *testing.T, filePath string) map[string]*models.TabEntry {
	tabsMap := make(map[string]*models.TabEntry)

	records, err := logic.GetCSVRecords(filePath, false)
	if err != nil {
		t.Fatalf("error occurred during the creation of tabs from a CSV: %s", err.Error())
	}

	for _, record := range records {
		tabID, _ := logic.StringToUUID(record[0])

		instrumentID, _ := strconv.ParseUint(record[1], 10, 64)
		difficultyID, _ := strconv.ParseUint(record[2], 10, 64)

		tab := &models.TabEntry{
			ID:           tabID,
			InstrumentID: uint(instrumentID),
			DifficultyID: uint(difficultyID),
			Description:  record[3],
		}

		tabsMap[tabID.String()] = tab
	}

	return tabsMap
}
