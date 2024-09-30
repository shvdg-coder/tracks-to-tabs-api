package environments

import (
	"github.com/shvdg-coder/base-logic/pkg/testable/database"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/services"
)

// DbEnvOps represents operations for a test environment, used during integration tests.
type DbEnvOps interface {
	database.ContainerOps
	services.CreateOps
	services.DropOps
	Setup()
	Breakdown()
}

// DbEnv is used to spin up a Database container for integration testing.
type DbEnv struct {
	database.ContainerOps
	services.CreateOps
	services.DropOps
}

// NewDbEnv creates a new instance of DbEnv.
func NewDbEnv(container database.ContainerOps, create services.CreateOps, drop services.DropOps) DbEnvOps {
	return &DbEnv{
		ContainerOps: container,
		CreateOps:    create,
		DropOps:      drop,
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
