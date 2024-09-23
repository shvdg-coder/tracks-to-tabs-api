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
	logic.DbOps
}

// API provides functionalities regarding the app.
type API struct {
	CreateOps
	DropOps
	DataOps
	DummyOps
	SeedOps
	logic.DbOps
}

// NewAPI instantiates a API.
func NewAPI(configPath string) (*API, error) {
	apiConfig, err := NewAPIConfig(configPath)
	if err != nil {
		return nil, err
	}

	dbConfig := apiConfig.Database
	database := logic.NewDbSvc(ValueDatabaseDriver, dbConfig.URL, logic.WithSSHTunnel(dbConfig.SSH), logic.WithConnection())

	svcManager := NewSvcManager(database)
	api := &API{
		CreateOps: NewCreateAPI(svcManager),
		DropOps:   NewDropAPI(svcManager),
		DataOps:   NewDataAPI(svcManager),
		DbOps:     database,
	}

	seeding := apiConfig.Seeding
	if seeding != nil {
		dummyAPI := NewDummyAPI(svcManager, seeding.Sources, seeding.Instruments, seeding.Difficulties)
		api.DummyOps = dummyAPI
		api.SeedOps = NewSeedingAPI(svcManager, seeding, dummyAPI)
	}

	return api, nil
}
