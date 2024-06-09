package main

import (
	"fmt"
	logic "github.com/shvdg-dev/base-logic/pkg"
	"github.com/shvdg-dev/tunes-to-tabs-api/internal"
	"github.com/shvdg-dev/tunes-to-tabs-api/pkg"
	"os"
)

const errMsg = "Failed to run app, expected 'create', 'purge', or 'seed'"

// main is the entry point of the application.
func main() {
	if len(os.Args) < 2 {
		printErrorAndExit()
	}
	URL := logic.GetEnvValueAsString(internal.KeyDatabaseURL)
	database := logic.NewDatabaseManager(internal.ValueDatabaseDriver, URL)
	api := pkg.NewAPI(database)

	arg := os.Args[1]
	handleArg(arg, api)
}

// handleArg handles the command line argument and performs the corresponding action.
func handleArg(arg string, api *pkg.API) {
	switch arg {
	case "create":
		internal.NewCreator(api).CreateTables()
	case "purge":
		internal.NewPurger(api).DropTables()
	case "seed":
		internal.NewSeeder(api).SeedTables()
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println(errMsg)
	os.Exit(1)
}
