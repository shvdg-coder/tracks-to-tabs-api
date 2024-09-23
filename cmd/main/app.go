package main

import (
	"fmt"
	"github.com/shvdg-coder/tracks-to-tabs-api/internal"
	"github.com/shvdg-coder/tracks-to-tabs-api/pkg"
	"log"
	"os"
)

var (
	api *pkg.API
)

// init instantiates all app requirements.
func init() {
	var err error
	api, err = pkg.NewAPI(internal.APIConfigPath)
	if err != nil {
		log.Fatal(err)
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
	case internal.CommandCreate:
		api.CreateAll()
	case internal.CommandDrop:
		api.DropAll()
	case internal.CommandSeed:
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
