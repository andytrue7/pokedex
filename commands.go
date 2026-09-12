package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

// catchThreshold controls how hard Pokémon are to catch: for each attempt
// we roll a number in [0, BaseExperience) and catch the Pokémon only if it
// lands at or below this threshold. Higher BaseExperience widens the roll
// range, so tougher Pokémon are less likely to land under the threshold
// and are correspondingly harder to catch.
const catchThreshold = 40

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
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
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

func commandCatch(config *ReplStateConfig, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}
	name := args[0]

	pokemon, err := config.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	// rand.Intn panics on n <= 0, and some Pokémon report a BaseExperience
	// of 0; treat those as trivially easy to catch instead.
	baseExperience := pokemon.BaseExperience
	if baseExperience <= 0 {
		baseExperience = 1
	}

	if rand.Intn(baseExperience) > catchThreshold {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	config.pokedex[name] = pokemon
	return nil
}

func commandInspect(config *ReplStateConfig, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}
	name := args[0]

	pokemon, ok := config.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(config *ReplStateConfig, args []string) error {
	names := make([]string, 0, len(config.pokedex))
	for name := range config.pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
