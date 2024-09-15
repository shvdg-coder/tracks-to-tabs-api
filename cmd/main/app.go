package main

import (
	"fmt"
	inter "github.com/shvdg-dev/tracks-to-tabs-api/internal"
	"github.com/shvdg-dev/tracks-to-tabs-api/pkg"
	"os"
)

var (
	api *pkg.API
)

// init instantiates all app requirements.
func init() {
	api = pkg.NewAPI()
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
		api.CreateAll()
	case inter.CommandDrop:
		api.DropAll()
	case inter.CommandSeed:
		api.Seed()
	default:
		printErrorAndExit()
	}
}

// printErrorAndExit prints an error message and exits the program with an exit code of 1.
func printErrorAndExit() {
	fmt.Println("Failed to run app, expected 'create', 'drop', or 'seed'")
	os.Exit(1)
}
