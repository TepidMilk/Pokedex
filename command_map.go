package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationsResp.Next
	cfg.Previous = locationsResp.Previous

	for _, area := range locationsResp.Results {
		fmt.Println(area.Name)
	}
	return nil
}
