package pokeapi

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// ListLocationAreas fetches a page of location areas. Pass an empty
// pageURL to get the first page; otherwise pass the Next or Previous URL
// returned by a prior call to page forwards/backwards.
func (c *Client) ListLocationAreas(pageURL string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != "" {
		url = pageURL
	}

	locationAreas := LocationAreas{}
	if err := c.get(url, &locationAreas); err != nil {
		return LocationAreas{}, err
	}

	return locationAreas, nil
}
