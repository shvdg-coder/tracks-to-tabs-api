package pkg

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
)

// APIOps represents the operations of the API.
type APIOps interface {
	CreateOps
	DropOps
	DataOps
	DummyOps
	SeedOps
}

// API provides functionalities regarding the app.
type API struct {
	CreateOps
	DropOps
	DataOps
	DummyOps
	SeedOps
}

// NewAPI instantiates a API.
func NewAPI(configPath string) (*API, error) {
	apiConfig, err := NewAPIConfig(configPath)
	if err != nil {
		return nil, err
	}

	svcManager := createServiceManager(apiConfig.Database)
	api := &API{
		CreateOps: NewCreateAPI(svcManager),
		DropOps:   NewDropAPI(svcManager),
		DataOps:   NewDataAPI(svcManager),
	}

	seeding := apiConfig.Seeding
	if seeding != nil {
		dummyAPI := NewDummyAPI(svcManager, seeding.Sources, seeding.Instruments, seeding.Difficulties)
		api.SeedOps = NewSeedingAPI(svcManager, seeding, dummyAPI)
	}

	return api, nil
}

// createServiceManager instantiates the service manager with the database.
func createServiceManager(dbConfig *DatabaseConfig) *SvcManager {
	database := logic.NewDbSvc(ValueDatabaseDriver, dbConfig.URL, logic.WithSSHTunnel(dbConfig.SSH), logic.WithConnection())
	return NewSvcManager(database)
}
