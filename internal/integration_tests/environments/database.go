package environments

import (
	"github.com/shvdg-dev/base-logic/pkg/testable/database"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
)

// DbEnvOperations represents operations for a test environment, used during integration tests.
type DbEnvOperations interface {
	database.ContainerOperations
	pkg.CreateOps
	pkg.DropOps
	Setup()
	Breakdown()
}

// DbEnv is used to spin up a Database container for integration testing.
type DbEnv struct {
	database.ContainerOperations
	pkg.CreateOps
	pkg.DropOps
}

// NewDbEnv creates a new instance of DbEnv.
func NewDbEnv(dbContainer database.ContainerOperations, create pkg.CreateOps, drop pkg.DropOps) DbEnvOperations {
	return &DbEnv{
		ContainerOperations: dbContainer,
		CreateOps:           create,
		DropOps:             drop,
	}
}

// Setup prepares the DbEnv.
func (t *DbEnv) Setup() {
	t.CreateAll()
}

// Breakdown cleans up and breaks down the DbEnv.
func (t *DbEnv) Breakdown() {
	t.DropAll()
	t.Disconnect()
	t.Teardown()
}
