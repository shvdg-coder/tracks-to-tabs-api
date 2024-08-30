package integration_tests

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"testing"
)

// ExpectedTab contains the data of what a models.Tab is expected to have.
type ExpectedTab struct {
	ID              string
	InstrumentID    uint
	DifficultyID    uint
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

	testFieldsOfTabs(t, actualTabs, createExpectedTabs())
}

// createExpectedTabs constructs and returns a slice of ExpectedTab objects for use in test cases.
func createExpectedTabs() []*ExpectedTab {
	return []*ExpectedTab{
		{ID: "3fa85f64-5717-4562-b3fc-2c963f66afa6", InstrumentID: 1, DifficultyID: 3, ReferencesCount: 2, ResourceCount: 1},
		{ID: "f79e3f20-a634-4c3e-90a9-70c5fe8b0195", InstrumentID: 1, DifficultyID: 2, ReferencesCount: 2, ResourceCount: 1},
		{ID: "6fa91502-efd5-4f52-9087-0b8bf7343f2b", InstrumentID: 1, DifficultyID: 2, ReferencesCount: 2, ResourceCount: 1},
		{ID: "337ab4e4-2c48-41f8-9131-9acf511d72a6", InstrumentID: 1, DifficultyID: 3, ReferencesCount: 2, ResourceCount: 1},
	}
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
	tabID, _ := logic.UUIDToString(actualTab.ID)
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
