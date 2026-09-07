// Package pokeapi is a thin client for the PokeAPI (https://pokeapi.co).
// Each endpoint the CLI needs gets its own file (e.g. location_areas.go,
// pokemon.go) with methods on Client; this file only holds the shared
// plumbing (the http.Client, base URL, and a generic GET+decode helper).
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// get issues a GET request against url and decodes the JSON response body
// into target, which must be a pointer.
func (c *Client) get(url string, target any) error {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	return json.Unmarshal(body, target)
}
