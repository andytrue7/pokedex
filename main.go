package main

import (
	"time"

	"github.com/andytrue7/pokedexcli/internal/pokeapi"
)

func main() {
	config := &ReplStateConfig{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	startRepl(config)
}
