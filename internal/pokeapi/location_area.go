package pokeapi

// LocationArea is the detail view of a single location area, e.g. the
// response from GET /location-area/{name}/. It's distinct from
// LocationAreas (plural), which is the paginated list of area names/URLs.
type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// GetLocationArea fetches the details of a single location area by name
// (e.g. "canalave-city-area"), including the Pokémon that can be
// encountered there.
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	locationArea := LocationArea{}
	if err := c.get(url, &locationArea); err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
