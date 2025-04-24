package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	dat, ok := c.cache.Get(url)
	if ok {
		locationsResp := Locations{}
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

	locationsResp := Locations{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil

}
