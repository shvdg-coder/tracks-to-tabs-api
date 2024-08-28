package main

import (
	"fmt"
	logic "github.com/shvdg-dev/base-logic/pkg"
	inter "github.com/shvdg-dev/tracks-to-tabs-api/internal"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"log"
	"os"
)

var (
	config    *inter.Config
	createAPI pkg.CreateOps
	dropAPI   pkg.DropOps
)

// init instantiates all app requirements.
func init() {
	config = initConfig()
	database := initDatabase()
	createAPI = pkg.NewCreateAPI(database)
	dropAPI = pkg.NewDropAPI(database)
}

// initConfig initializes the application configuration.
func initConfig() *inter.Config {
	conf, err := inter.NewConfig(inter.PathConfig)
	if err != nil {
		log.Fatalf("Could not load config")
	}
	return conf
}

// initDatabase initializes the database manager.
func initDatabase() logic.DbOperations {
	URL := logic.GetEnvValueAsString(inter.KeyDatabaseURL)
	database := logic.NewDbService(inter.ValueDatabaseDriver, URL)
	return database
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
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println("Failed to run app, expected 'create', 'purge', or 'seed'")
	os.Exit(1)
}
