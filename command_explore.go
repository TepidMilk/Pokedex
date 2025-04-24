package main

import (
	"errors"
	"fmt"
)

func commandExplore(location string, cfg *Config) error {
	if location == "" {
		return errors.New("invalid location area")
	}
	exploreResp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	cfg.current = &location

	fmt.Printf("Exploring %s...\n", location)
	fmt.Print("Found Pokemon:\n")
	for _, pokemonEncounter := range exploreResp.PokemonEncounters {
		fmt.Println(" - ", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
