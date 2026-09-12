package pokeapi

// Pokemon is the detail view of a single Pokémon, e.g. the response from
// GET /pokemon/{name}/.
type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// GetPokemon fetches the details of a single Pokémon by name (e.g. "pikachu").
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	pokemon := Pokemon{}
	if err := c.get(url, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
