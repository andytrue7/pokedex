package main

import (
	"time"

	"github.com/andytrue7/pokedexcli/internal/pokeapi"
)

const (
	httpTimeout   = 5 * time.Second
	cacheInterval = 5 * time.Minute
)

func main() {
	config := &ReplStateConfig{
		commands:      getCommands(),
		pokeapiClient: pokeapi.NewClient(httpTimeout, cacheInterval),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(config)
}
