package pkg

import (
	logic "github.com/shvdg-coder/base-logic/pkg"
	inter "github.com/shvdg-dev/tracks-to-tabs-api/internal"
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
	dummyAPI := NewDummyAPI(database, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)

	return &API{
		CreateOps: NewCreateAPI(database),
		DropOps:   NewDropAPI(database),
		DataOps:   NewDataAPI(database),
		DummyOps:  dummyAPI,
		SeedOps:   NewSeedingAPI(database, seedConfig, dummyAPI),
	}
}

// initSeedConfig instantiates the config for seeding.
func initSeedConfig() *models.SeedConfig {
	conf, err := models.NewSeedConfig(inter.PathSeedConfig)
	if err != nil {
		log.Fatalf("Could not load seed config")
	}
	return conf
}

// initDatabase instantiates the database.
func initDatabase() logic.DbOperations {
	dbURL := logic.GetEnvValueAsString(inter.KeyDatabaseURL)
	sshConfig := createSSHConfig()

	database := logic.NewDbService(inter.ValueDatabaseDriver, dbURL, logic.WithSSHTunnel(sshConfig), logic.WithConnection())
	return database
}

// createSSHConfig
func createSSHConfig() *logic.SSHConfig {
	return &logic.SSHConfig{
		User:        logic.GetEnvValueAsString(inter.KeySshUser),
		Password:    logic.GetEnvValueAsString(inter.KeySshPassword),
		Server:      logic.GetEnvValueAsString(inter.KeySshServer),
		Destination: logic.GetEnvValueAsString(inter.KeySshDestination),
		LocalPort:   logic.GetEnvValueAsString(inter.KeySshLocalPort),
	}
}
