package testable

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/docker/go-connections/nat"
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

	return &TestDatabase{
		Resource:      container,
		DB:            db,
		Host:          host,
		Port:          port.Port(),
		WritePassword: PostgresPassword,
		WriteUser:     PostgresUser,
		WriteDBName:   PostgresDB,
	}, nil
}

// NewPostgresContainer creates and starts a Postgres Docker container.
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

// GetContainerInfo gets the IP and the port map from a running Docker container.
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

// CreateURL creates the URL to the database
func (t *TestDatabase) CreateURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		t.Host,
		t.Port,
		t.WriteUser,
		t.WritePassword,
		t.WriteDBName)
}

// Teardown brings the database down.
func (t *TestDatabase) Teardown() error {
	return t.Resource.Terminate(context.Background())
}
