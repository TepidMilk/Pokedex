package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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
	var locationsResp Locations
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locationsResp); err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
