package testable

import (
	"github.com/shvdg-dev/base-logic/pkg/testable/database"
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
)

// TestEnvOperations represents operations for a test environment, used during integration tests.
type TestEnvOperations interface {
	database.ContainerOperations
	inl.CreateOperations
	inl.DropOperations
	Setup()
	Breakdown()
}

// TestEnv is used to spin up a database container for integration testing.
type TestEnv struct {
	database.ContainerOperations
	inl.CreateOperations
	inl.DropOperations
}

// NewTestEnv creates a new instance of TestEnv.
func NewTestEnv(dbContainer database.ContainerOperations, create inl.CreateOperations, drop inl.DropOperations) TestEnvOperations {
	return &TestEnv{
		ContainerOperations: dbContainer,
		CreateOperations:    create,
		DropOperations:      drop,
	}
}

// Setup prepares the TestEnv.
func (t *TestEnv) Setup() {
	t.CreateAll()
}

// Breakdown cleans up and breaks down the TestEnv.
func (t *TestEnv) Breakdown() {
	t.DropAll()
	t.Disconnect()
	t.Teardown()
}
