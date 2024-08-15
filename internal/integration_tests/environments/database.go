package environments

import (
	"github.com/shvdg-dev/base-logic/pkg/testable/database"
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
)

// DbEnvOperations represents operations for a test environment, used during integration tests.
type DbEnvOperations interface {
	database.ContainerOperations
	inl.CreateOperations
	inl.DropOperations
	Setup()
	Breakdown()
}

// DbEnv is used to spin up a Database container for integration testing.
type DbEnv struct {
	database.ContainerOperations
	inl.CreateOperations
	inl.DropOperations
}

// NewDbEnv creates a new instance of DbEnv.
func NewDbEnv(dbContainer database.ContainerOperations, create inl.CreateOperations, drop inl.DropOperations) DbEnvOperations {
	return &DbEnv{
		ContainerOperations: dbContainer,
		CreateOperations:    create,
		DropOperations:      drop,
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
