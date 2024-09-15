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
	config, err := NewAPIConfig(configPath)
	if err != nil {
		return nil, err
	}

	database := createDatabase(config.Database)
	svcManager := NewSvcManager(database)
	seeding := config.Seeding
	dummyAPI := NewDummyAPI(svcManager, seeding.Sources, seeding.Instruments, seeding.Difficulties)

	return &API{
		CreateOps: NewCreateAPI(svcManager),
		DropOps:   NewDropAPI(svcManager),
		DataOps:   NewDataAPI(svcManager),
		DummyOps:  dummyAPI,
		SeedOps:   NewSeedingAPI(svcManager, seeding, dummyAPI),
	}, nil
}

// createDatabase instantiates the database.
func createDatabase(dbConfig *DatabaseConfig) logic.DbOperations {
	return logic.NewDbService(
		ValueDatabaseDriver, dbConfig.URL,
		logic.WithSSHTunnel(dbConfig.SSH), logic.WithConnection())
}
