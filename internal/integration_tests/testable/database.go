package testable

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
	logic "github.com/shvdg-dev/base-logic/pkg"
	inl "github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// TestDatabase is used to spin up a database for integration testing.
type TestDatabase struct {
	testcontainers.Container
	*sql.DB
	inl.CreateOperations
	inl.DropOperations
}

// NewTestDatabase creates a new instance of TestDatabase.
func NewTestDatabase() (*TestDatabase, error) {
	ctx := context.Background()

	container, err := NewPostgresContainer(ctx)
	if err != nil {
		return nil, err
	}

	host, port, err := GetContainerInfo(ctx, container)
	if err != nil {
		return nil, err
	}

	db, err := NewPostgresConnection(host, port)
	if err != nil {
		return nil, err
	}

	databaseManager := logic.NewDatabaseManager("postgres", CreateURL(host, port.Port()))
	tableService := inl.NewTableService(databaseManager)

	return &TestDatabase{
		Container:        container,
		DB:               db,
		CreateOperations: inl.NewCreateService(tableService),
		DropOperations:   inl.NewDropService(tableService),
	}, nil
}

// NewPostgresContainer sets up a new Postgres container in Docker and returns the container instance.
func NewPostgresContainer(ctx context.Context) (testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        PostgresImage,
		ExposedPorts: []string{PostgresExposedPorts},
		WaitingFor:   wait.ForListeningPort(PostgresExposedPorts),
		Env: map[string]string{
			"POSTGRES_PASSWORD": PostgresPassword,
			"POSTGRES_USER":     PostgresUser,
			"POSTGRES_DB":       PostgresDB,
		},
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

// GetContainerInfo fetches and returns the host IP and mapped port from a running Docker container.
func GetContainerInfo(ctx context.Context, container testcontainers.Container) (ip string, port nat.Port, err error) {
	ip, err = container.Host(ctx)
	if err != nil {
		return "", "", err
	}

	port, err = container.MappedPort(ctx, PostgresPort)
	if err != nil {
		return "", "", err
	}

	return ip, port, nil
}

// NewPostgresConnection creates a new DB connection to a Postgres service.
func NewPostgresConnection(host string, port nat.Port) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port.Port(), PostgresUser, PostgresPassword, PostgresDB)
	return sql.Open("postgres", dataSource)
}

// CreateURL constructs and returns the database connection URL using the constants and provided host and port.
func CreateURL(host, port string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		PostgresUser,
		PostgresPassword,
		PostgresDB)
}

// Teardown destroys the database container.
func (t *TestDatabase) Teardown() error {
	return t.Terminate(context.Background())
}
