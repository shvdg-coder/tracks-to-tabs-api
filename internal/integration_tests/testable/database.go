package testable

import (
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
)

// TestDatabase is used to spin up a database for integration testing.
type TestDatabase struct {
	inl.CreateOperations
	inl.DropOperations
}
