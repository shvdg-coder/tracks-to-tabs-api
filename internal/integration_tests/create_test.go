package integration_tests

import (
	tstdb "github.com/shvdg-dev/base-logic/pkg/testable/database"
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
	tstabl "github.com/shvdg-dev/tunes-to-tabs-api/internal/integration_tests/testable"
	"testing"
)

func TestEnvSetup(t *testing.T) {
	dbContainerService := tstdb.NewContainerService()
	config := tstdb.NewPostgresContainerConfig()
	dbContainer, err := dbContainerService.CreateContainer(config)
	if err != nil {
		t.Fatal(err)
	}

	tablesService := inl.NewTableService(dbContainer)
	creatorService := inl.NewCreateService(tablesService)
	dropService := inl.NewDropService(tablesService)

	testEnv := tstabl.NewTestEnv(dbContainer, creatorService, dropService)

	defer testEnv.Breakdown()

	err = testEnv.Ping()
	if err != nil {
		t.Fatal(err)
	}

	testEnv.Setup()
}
