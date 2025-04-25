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

	//list of current pokemon
	currentPokemon := []string{}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Print("Found Pokemon:\n")
	for _, pokemonEncounter := range exploreResp.PokemonEncounters {
		fmt.Println(" - ", pokemonEncounter.Pokemon.Name)
		currentPokemon = append(currentPokemon, pokemonEncounter.Pokemon.Name)
	}

	//configure current pokemon
	cfg.current = currentPokemon

	return nil
}
