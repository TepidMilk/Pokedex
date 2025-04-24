package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	var locationsResp Locations
	if pageURL != nil {
		url = *pageURL
	}
	dat, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return Locations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}
	c.cache.Add(url, body)

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locationsResp); err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
