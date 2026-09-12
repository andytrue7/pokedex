# Pokédex CLI

A terminal Pokédex: an interactive REPL, written in Go, that talks to the
[PokéAPI](https://pokeapi.co) to browse location areas, explore the
Pokémon found there, and catch and inspect them.

## Motivation

This is a learning project for practicing Go fundamentals — REPLs,
structuring an HTTP API client as its own package, JSON decoding, and a
simple in-process TTL cache — by building something with a real, freely
available API (PokeAPI) instead of a toy dataset. It grew feature by
feature (map, explore, catch, inspect, pokedex) the way a real CLI tool
would.

## Quick Start

Requirements:

- Go 1.26 or later
- Internet access (all data comes from the public PokeAPI at
  `https://pokeapi.co/api/v2`)

Run it:

```sh
go run .
```

This drops you into the REPL:

```
Pokedex > 
```

Type a command and press enter. Type `exit` to quit.

## Usage

### Commands

| Command | Arguments | Description |
|---|---|---|
| `help` | — | Lists all available commands |
| `map` | — | Shows the next page of 20 location areas |
| `mapb` | — | Shows the previous page of location areas |
| `explore <location-area>` | area name | Lists the Pokémon that can be encountered in that area |
| `catch <pokemon>` | Pokémon name | Attempts to catch the Pokémon; success depends on its base experience (tougher Pokémon are harder to catch) |
| `inspect <pokemon>` | Pokémon name | Prints the name, height, weight, stats, and type(s) of a Pokémon you've already caught |
| `pokedex` | — | Lists the names of every Pokémon you've caught so far, alphabetically |
| `exit` | — | Quits the Pokedex |

### Example session

```
Pokedex > map
Map:

canalave-city-area
eterna-city-area
...

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - staryu
 - magikarp
...

Pokedex > catch pidgey
Throwing a Pokeball at pidgey...
pidgey was caught!

Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying

Pokedex > pokedex
Your Pokedex:
 - pidgey

Pokedex > exit
Closing the Pokedex... Goodbye!
```

### Features

- **Paginated location browsing** (`map` / `mapb`) — walks forward and
  backward through PokeAPI's paginated `location-area` list, remembering
  the current page's `next`/`previous` URLs between commands.
- **Location exploration** (`explore`) — looks up a single location area
  by name and lists every Pokémon known to appear there.
- **Catching** (`catch`) — simulates throwing a Pokéball; the catch chance
  is inversely related to the Pokémon's base experience, so common
  Pokémon are caught reliably and legendaries are much harder. Caught
  Pokémon are kept in an in-memory Pokédex for the session.
- **Inspecting** (`inspect`) — shows full details (height, weight, base
  stats, types) for any Pokémon already in your Pokédex.
- **In-memory Pokédex** (`pokedex`) — lists everything you've caught.
- **Response caching** — all PokeAPI requests are cached in-process for a
  few minutes, so re-fetching a page or Pokémon (e.g. paging back and
  forth, or re-catching the same species) doesn't hit the network again.

## Contributing

This started as a personal practice project, but issues and PRs are
welcome.

Project layout, if you're getting oriented:

```
main.go                       entry point: builds the config and starts the REPL
repl.go                       the read-eval-print loop and input parsing
commands.go                   command registry and all command implementations
config.go                     REPL session state (pagination cursors, caught Pokédex)
internal/pokeapi/             PokeAPI HTTP client, one file per endpoint
  client.go                     shared http.Client + cached GET/decode helper
  location_areas.go             GET /location-area (paginated list)
  location_area.go              GET /location-area/{name} (single area detail)
  pokemon.go                    GET /pokemon/{name} (single Pokémon detail)
internal/pokecache/           short-lived TTL cache used by the PokeAPI client
```

Adding a new endpoint means adding one file to `internal/pokeapi/` with a
response struct and a method on `*Client` that calls the shared `get()`
helper — you get caching for free.

Before opening a PR:

```sh
gofmt -l .        # should print nothing
go vet ./...
go test -race ./...
```
