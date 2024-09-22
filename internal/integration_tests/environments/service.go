package environments

import (
	tstdb "github.com/shvdg-coder/base-logic/pkg/testable/database"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg"
)

// EnvManagement represents the operations related to managing environments.
type EnvManagement interface {
	CreatePostgresEnv() (DbEnvOps, error)
	CreateDbEnv(config *tstdb.ContainerConfig) (DbEnvOps, error)
}

// EnvSvc is responsible for managing different environments required for integration testing.
type EnvSvc struct {
	EnvManagement
	Database tstdb.ContainerManagement
}

// NewEnvSvc instantiates a new EnvSvc.
func NewEnvSvc() EnvManagement {
	return &EnvSvc{
		Database: tstdb.NewContainerSvc(),
	}
}

// CreatePostgresEnv creates a Database environment for Postgres, with default configurations.
func (s *EnvSvc) CreatePostgresEnv() (DbEnvOps, error) {
	return s.CreateDbEnv(tstdb.NewPostgresContainerConfig())
}

// CreateDbEnv creates a Database environment.
func (s *EnvSvc) CreateDbEnv(config *tstdb.ContainerConfig) (DbEnvOps, error) {
	dbContainer, err := s.Database.CreateContainer(config)
	svcManager := pkg.NewSvcManager(dbContainer)
	if err != nil {
		return nil, err
	}

	creatorService := pkg.NewCreateAPI(svcManager)
	dropService := pkg.NewDropAPI(svcManager)
	dbEnv := NewDbEnv(dbContainer, creatorService, dropService)

	return dbEnv, nil
}
