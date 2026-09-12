package main

import "github.com/andytrue7/pokedexcli/internal/pokeapi"

type ReplStateConfig struct {
	commands            map[string]cliCommand
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL string
	prevLocationAreaURL string
	pokedex             map[string]pokeapi.Pokemon
}

// storeLocationAreaPage remembers the Next/Previous URLs of a fetched page
// so the following map/mapb call knows where to go. PokeAPI returns null
// for Next/Previous at the ends of the list, which we normalize to "".
func (c *ReplStateConfig) storeLocationAreaPage(page pokeapi.LocationAreas) {
	c.nextLocationAreaURL = ""
	if page.Next != nil {
		c.nextLocationAreaURL = *page.Next
	}

	c.prevLocationAreaURL = ""
	if page.Previous != nil {
		c.prevLocationAreaURL = *page.Previous
	}
}
