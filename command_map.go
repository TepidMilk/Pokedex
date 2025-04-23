package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMap(c *Config) error {
	if c.Next == "" {
		c.Next = "https://pokeapi.co/api/v2/location-area/"
		c.Previous = ""
	}
	url := c.Next
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var location Locations
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&location); err != nil {
		return err
	}

	for _, area := range location.Results {
		fmt.Println(area.Name)
	}
	c.Next = location.Next
	c.Previous = location.Previous
	return nil
}

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
