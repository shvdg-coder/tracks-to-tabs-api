package pkg

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/constants"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/models"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg/services"
)

// APIOps represents the operations of the API.
type APIOps interface {
	services.CreateOps
	services.DropOps
	services.DataOps
	services.DummyOps
	services.SeedOps
	logic.DbOps
}

// API provides functionalities regarding the app.
type API struct {
	Config *models.APIConfig
	services.CreateOps
	services.DropOps
	services.DataOps
	services.DummyOps
	services.SeedOps
	logic.DbOps
}

// NewAPI instantiates a API.
func NewAPI(configPath string) (*API, error) {
	apiConfig, err := models.NewAPIConfig(configPath)
	if err != nil {
		return nil, err
	}

	dbConfig := apiConfig.Database
	database := logic.NewDbSvc(constants.ValueDatabaseDriver, dbConfig.URL, logic.WithSSHTunnel(dbConfig.SSH), logic.WithConnection())

	svcManager := services.NewSvcManager(database)
	api := &API{
		Config:    apiConfig,
		CreateOps: services.NewCreateSvc(svcManager),
		DropOps:   services.NewDropSvc(svcManager),
		DataOps:   services.NewDataSvc(svcManager),
		DbOps:     database,
	}

	seeding := apiConfig.Seeding
	if seeding != nil {
		dummySvc := services.NewDummySvc(svcManager, seeding.Sources, seeding.Instruments, seeding.Difficulties)
		api.DummyOps = dummySvc
		api.SeedOps = services.NewSeedSvc(svcManager, seeding, dummySvc)
	}

	return api, nil
}
