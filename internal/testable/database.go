package testable

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestDatabase holds data used to spin up a database for integration testing.
type TestDatabase struct {
	Resource testcontainers.Container
	*sql.DB
	Host                                  string
	Port                                  string
	WritePassword, WriteUser, WriteDBName string
}

// NewTestDatabase creates a new instance of TestDatabase.
func NewTestDatabase() (*TestDatabase, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_PASSWORD": "docker",
			"POSTGRES_USER":     "docker",
			"POSTGRES_DB":       "test",
		},
	}

	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := postgres.Host(ctx)
	if err != nil {
		return nil, err
	}

	port, err := postgres.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", ip, port.Port(), "docker", "docker", "test")
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	return &TestDatabase{
		Resource:      postgres,
		DB:            db,
		Host:          ip,
		Port:          port.Port(),
		WritePassword: "docker",
		WriteUser:     "docker",
		WriteDBName:   "test",
	}, nil
}

// Teardown brings the database down.
func (t *TestDatabase) Teardown() error {
	return t.Resource.Terminate(context.Background())
}
