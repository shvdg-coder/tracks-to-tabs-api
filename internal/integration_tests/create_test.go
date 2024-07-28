package integration_tests

import (
	"github.com/shvdg-dev/tunes-to-tabs-api/internal/integration_tests/testable"
	"testing"
)

func TestTableCreation(t *testing.T) {
	db, err := testable.NewTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Teardown()

	db.CreateAll()
	db.DropAll()
}
