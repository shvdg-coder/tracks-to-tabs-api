package integration_tests

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"strconv"
	"testing"
)

// ExpectedTab contains the data of what a models.Tab is expected to have.
type ExpectedTab struct {
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

	expectedTabsMap := createExpectedTabs(t)

	tabIDStrings, err := logic.GetCSVColumnValues(tabsCSV, tabsColumnID)
	if err != nil {
		t.Fatal(err)
	}

	tabIDs, err := logic.StringsToUUIDs(tabIDStrings...)
	if err != nil {
		t.Fatal(err)
	}

	mapper := mappers.NewTabSvc()
	api := pkg.NewDataAPI(dbEnv)

	// Execute
	actualTabs, err := api.GetTabs(tabIDs...)
	if err != nil {
		t.Fatalf("error occurred during retrieval of tab: %s", err.Error())
	}

	// Test
	if len(actualTabs) == 0 || len(tabIDs) == 0 {
		t.Errorf("expected more than 0 tabs (CSV: %d, API: %d)", len(tabIDs), len(actualTabs))
	}

	if len(actualTabs) != len(tabIDs) {
		t.Errorf("expected %d tabs, got %d", len(tabIDs), len(actualTabs))
	}

	testFieldsOfTabs(t, mapper.TabsToMap(actualTabs), expectedTabsMap)
}

// testFieldsOfTabs tests the fields of multiple tab objects by comparing the actual tabs to the expected ones.
func testFieldsOfTabs(t *testing.T, actualTabsMap map[uuid.UUID]*models.Tab, expectedTabsMap map[uuid.UUID]*ExpectedTab) {
	for id := range actualTabsMap {
		actualTab := actualTabsMap[id]
		expectedTab, ok := expectedTabsMap[id]
		if !ok {
			t.Fatalf("ID %s does not exist in 'expected tabs' map", id)
		} else {
			testFieldsOfTab(t, actualTab, expectedTab)
		}
	}
}

// testFieldsOfTab tests the fields of a single tab object by comparing the actual tab to the expected one.
func testFieldsOfTab(t *testing.T, actualTab *models.Tab, expectedTab *ExpectedTab) {
	// Check ID
	if actualTab.ID != expectedTab.ID {
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

// createExpectedTabs constructs and returns a map of ExpectedTab's for use in test cases.
func createExpectedTabs(t *testing.T) map[uuid.UUID]*ExpectedTab {
	expectedTabs := make(map[uuid.UUID]*ExpectedTab)

	tabsMap := createTabsFromCSV(t, tabsCSV)
	for _, tab := range tabsMap {
		expectedTab := &ExpectedTab{
			TabEntry:        tab,
			ReferencesCount: 2,
			ResourceCount:   1,
		}
		expectedTabs[tab.ID] = expectedTab
	}

	return expectedTabs
}

// createTabsFromCSV creates a map of tabs where the key is the ID and the value a models.TabEntry.
func createTabsFromCSV(t *testing.T, filePath string) map[uuid.UUID]*models.TabEntry {
	tabsMap := make(map[uuid.UUID]*models.TabEntry)

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

		tabsMap[tabID] = tab
	}

	return tabsMap
}
