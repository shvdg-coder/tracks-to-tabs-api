package main

import (
	"fmt"
	logic "github.com/shvdg-dev/base-logic/pkg"
	inter "github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"os"
)

const errMsg = "Failed to run app, expected 'create', 'purge', or 'seed'"

// main is the entry point of the application.
func main() {
	URL := logic.GetEnvValueAsString(inter.KeyDatabaseURL)
	database := logic.NewDatabaseManager(inter.ValueDatabaseDriver, URL)
	api := pkg.NewAPI(database)

	handleArgs(os.Args[1:], api)
}

// handleArgs handles each argument individually.
func handleArgs(args []string, api *pkg.API) {
	for _, arg := range args {
		handleArg(arg, api)
	}
}

// handleArgs handles the command line argument and performs the corresponding action.
func handleArg(arg string, api *pkg.API) {
	switch arg {
	case inter.CommandCreate:
		inter.NewCreator(api).CreateTables()
	case inter.CommandPurge:
		inter.NewPurger(api).DropTables()
	case inter.CommandSeed:
		inter.NewSeeder(api).SeedTables()
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println(errMsg)
	os.Exit(1)
}
