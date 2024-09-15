package main

import (
	"fmt"
	logic "github.com/shvdg-coder/base-logic/pkg"
	inter "github.com/shvdg-dev/tracks-to-tabs-api/internal"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg/models"
	"log"
	"os"
)

var (
	createAPI pkg.CreateOps
	dropAPI   pkg.DropOps
	seedAPI   pkg.SeedOps
)

// init instantiates all app requirements.
func init() {
	seedConfig := initSeedConfig()
	database := initDatabase()
	createAPI = pkg.NewCreateAPI(database)
	dropAPI = pkg.NewDropAPI(database)
	dummyAPI := pkg.NewDummyAPI(database, seedConfig.Sources, seedConfig.Instruments, seedConfig.Difficulties)
	seedAPI = pkg.NewSeedingAPI(database, seedConfig, dummyAPI)
}

// initSeedConfig initializes the application seeding configuration.
func initSeedConfig() *models.SeedConfig {
	conf, err := models.NewSeedConfig(inter.PathSeedConfig)
	if err != nil {
		log.Fatalf("Could not load seed config")
	}
	return conf
}

// initDatabase initializes the database manager.
func initDatabase() logic.DbOperations {
	dbURL := logic.GetEnvValueAsString(inter.KeyDatabaseURL)
	sshConfig := createSSHConfig()

	database := logic.NewDbService(inter.ValueDatabaseDriver, dbURL, logic.WithSSHTunnel(sshConfig), logic.WithConnection())
	return database
}

// CreateSSHConfig creates a new SSH client config with values from environment variables.
func createSSHConfig() *logic.SSHConfig {
	return &logic.SSHConfig{
		User:        logic.GetEnvValueAsString(inter.KeySshUser),
		Password:    logic.GetEnvValueAsString(inter.KeySshPassword),
		Server:      logic.GetEnvValueAsString(inter.KeySshServer),
		Destination: logic.GetEnvValueAsString(inter.KeySshDestination),
		LocalPort:   logic.GetEnvValueAsString(inter.KeySshLocalPort),
	}
}

// main is the entry point of the application.
func main() {
	handleArgs(os.Args[1:])
}

// handleArgs handles each argument individually.
func handleArgs(args []string) {
	for _, arg := range args {
		handleArg(arg)
	}
}

// handleArgs handles the command line argument and performs the corresponding action.
func handleArg(arg string) {
	switch arg {
	case inter.CommandCreate:
		createAPI.CreateAll()
	case inter.CommandPurge:
		dropAPI.DropAll()
	case inter.CommandSeed:
		seedAPI.Seed()
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println("Failed to run app, expected 'create', 'purge', or 'seed'")
	os.Exit(1)
}
