package testable

import (
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
)

// TestDbOps represents operations for a test database, used during integration tests.
type TestDbOps interface {
	inl.CreateOperations
	inl.DropOperations
	Setup()
	Teardown()
}

// TestDb is used to spin up a database for integration testing.
type TestDb struct {
	inl.CreateOperations
	inl.DropOperations
}

// NewTestDb creates a new instance of TestDb.
func NewTestDb(create inl.CreateOperations, drop inl.DropOperations) TestDbOps {
	return &TestDb{
		CreateOperations: create,
		DropOperations:   drop,
	}
}

// Setup prepares the TestDb.
func (t *TestDb) Setup() {
	t.CreateAll()
}

// Teardown cleans up and breaks down the TestDb.
func (t *TestDb) Teardown() {
	t.DropAll()
}
