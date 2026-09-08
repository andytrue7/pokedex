package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*ReplStateConfig, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the map backwards",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area",
			callback:    commandExplore,
		},
	}
}

func commandExit(config *ReplStateConfig, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *ReplStateConfig, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(config *ReplStateConfig, args []string) error {
	fmt.Println("Map:")
	fmt.Println("")

	locationAreas, err := config.pokeapiClient.ListLocationAreas(config.nextLocationAreaURL)
	if err != nil {
		return err
	}

	config.storeLocationAreaPage(locationAreas)

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapBack(config *ReplStateConfig, args []string) error {
	fmt.Println("Map Back:")
	fmt.Println("")

	if config.prevLocationAreaURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := config.pokeapiClient.ListLocationAreas(config.prevLocationAreaURL)
	if err != nil {
		return err
	}

	config.storeLocationAreaPage(locationAreas)

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandExplore(config *ReplStateConfig, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-area-name>")
	}
	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)
	fmt.Println("Found Pokemon:")

	locationArea, err := config.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
