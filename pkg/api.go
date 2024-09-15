package pkg

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"log"
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
func NewAPI() *API {
	seedConfig := initSeedConfig()
	database := initDatabase()
	svcManager := NewSvcManager(database)
	dummyAPI := NewDummyAPI(database, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)

	return &API{
		CreateOps: NewCreateAPI(svcManager),
		DropOps:   NewDropAPI(svcManager),
		DataOps:   NewDataAPI(svcManager),
		DummyOps:  dummyAPI,
		SeedOps:   NewSeedingAPI(svcManager, seedConfig, dummyAPI),
	}
}

// initSeedConfig instantiates the config for seeding.
func initSeedConfig() *models.SeedConfig {
	conf, err := models.NewSeedConfig(PathSeedConfig)
	if err != nil {
		log.Fatalf("Could not load seed config")
	}
	return conf
}

// initDatabase instantiates the database.
func initDatabase() logic.DbOperations {
	dbURL := logic.GetEnvValueAsString(KeyDatabaseURL)
	sshConfig := createSSHConfig()

	database := logic.NewDbService(ValueDatabaseDriver, dbURL, logic.WithSSHTunnel(sshConfig), logic.WithConnection())
	return database
}

// createSSHConfig
func createSSHConfig() *logic.SSHConfig {
	return &logic.SSHConfig{
		User:        logic.GetEnvValueAsString(KeySshUser),
		Password:    logic.GetEnvValueAsString(KeySshPassword),
		Server:      logic.GetEnvValueAsString(KeySshServer),
		Destination: logic.GetEnvValueAsString(KeySshDestination),
		LocalPort:   logic.GetEnvValueAsString(KeySshLocalPort),
	}
}
