package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) PokemonEncounter(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	dat, ok := c.cache.Get(url)
	if ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(dat, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, body)

	pokemonResp := Pokemon{}
	err = json.Unmarshal(body, &pokemonResp)
	if err != nil {
		return Pokemon{}, errors.New("invalid pokemon name")
	}

	return pokemonResp, nil
}
