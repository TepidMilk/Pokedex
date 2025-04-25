package main

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

func commandCatch(pokemon string, cfg *Config) error {
	//Once there is a valid pokemon retrieve pokemon data
	pokemonResp, err := cfg.pokeapiClient.PokemonEncounter(pokemon)
	if err != nil {
		return err
	}

	if !slices.Contains(cfg.current, pokemon) {
		return errors.New("pokemon does not exist in current explored area")
	}

	baseEXP := pokemonResp.BaseExperience

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	//logic to catch pokemon
	if !catchPokemon(baseEXP) {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		fmt.Printf("%s was caught!\n", pokemon)
	}

	return nil
}

func catchPokemon(baseEXP int) bool {
	M := rand.Intn(255) //Catch value
	health := rand.Intn(baseEXP)
	f := (baseEXP * 255 * 4) / (health * 12)
	return f >= M
}
