package internal

import (
	"fmt"
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

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Host,
		db.Port,
		db.WriteUser,
		db.WritePassword,
		db.WriteDBName)

	// TODO: add a way to stop monitoring for the database manager
	database := logic.NewDatabaseManager("postgres", dataSource)
	tables := NewTableService(database)
	create := NewCreateService(tables)
	create.CreateAll()
}
