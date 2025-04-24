package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (exploreDat, error) {
	url := baseURL + "/location-area/" + location
	dat, ok := c.cache.Get(url)
	if ok {
		exploreResp := exploreDat{}
		err := json.Unmarshal(dat, &exploreResp)
		if err != nil {
			return exploreDat{}, err
		}
		return exploreResp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return exploreDat{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return exploreDat{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return exploreDat{}, err
	}
	c.cache.Add(url, body)

	exploreResp := exploreDat{}
	err = json.Unmarshal(body, &exploreResp)
	if err != nil {
		return exploreDat{}, err
	}

	return exploreResp, nil
}
