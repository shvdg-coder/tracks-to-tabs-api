package internal

import (
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal/testable"
	"testing"
)

func TestTableCreation(t *testing.T) {
	db, err := testable.NewTestDatabase()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Teardown()

	database := logic.NewDatabaseManager("postgres", db.CreateURL())
	tables := NewTableService(database)
	create := NewCreateService(tables)
	create.CreateAll()
}
