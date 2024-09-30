package dataAPI

import (
	"github.com/google/uuid"
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/mappers"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/services"
	"testing"
)

// TestGetTabs tests whether tabs can be inserted and retrieved cascading.
func TestGetTabs(t *testing.T) {
	dbEnv := createDefaultDbEnv(t)
	svcManager := services.NewSvcManager(dbEnv)
	defer dbEnv.Breakdown()

	// Prepare
	seed(t, dbEnv, minConfigPath)
	insertCSVFiles(t, dbEnv)

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
	api := services.NewDataSvc(svcManager)

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
}
